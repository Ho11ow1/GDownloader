package Services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"path"
	"strings"
	"sync"

	"GDownloader/Common"
	"GDownloader/Models"
	"GDownloader/Utils"
)

func (this FilesterService) HandleDownload(pageURL string) {
    //
    if this.Base.IsAlbum(Common.Filester, pageURL) {
        urls, err := this.GetAlbumURLs(pageURL)
        if err != nil {
            Utils.Logger.LogError(fmt.Sprintf("Failed to get album URLs: %s", err))
            return
        }

        semaphore := make(chan struct{}, Common.AppDefs.MaxConcurrent)
        var wg sync.WaitGroup

        for _, slug := range urls {
            wg.Add(1)
            semaphore <- struct{}{}

            go func(s Models.FilesterSlug) {
                defer wg.Done()
                defer func() { <-semaphore }()

                _, downloadURL, err := this.GetFileInfo(s.URL)
                if err != nil {
                    Utils.Logger.LogError(fmt.Sprintf("Failed to get download URL: %s", err))
                    return
                }

                err = this.Base.Download(this.Client, s.Filename, s.URL, downloadURL, pageURL)
                if err != nil {
                    Utils.Logger.LogError(fmt.Sprintf("Download failed: %s", err))
                }
            }(slug)
        }

        wg.Wait()

    } else {
        filename, downloadURL, err := this.GetFileInfo(pageURL)
        if err != nil {
            Utils.Logger.LogError(fmt.Sprintf("Failed to get file info: %s", err))
            return
        }

        err = this.Base.Download(this.Client, filename, pageURL, downloadURL, pageURL);
        if err != nil {
            Utils.Logger.LogError(fmt.Sprintf("Download failed: %s", err))
        }
    }
}

func (this FilesterService) GetAlbumURLs(pageURL string) ([]Models.FilesterSlug, error) {
    //
    body, err := this.Client.Get(strings.Split(pageURL, `?`)[0] + `?page=1`)
    if err != nil {
        return nil, fmt.Errorf(`Failed to fetch album page 1: %w`, err)
    }

    return this.ParseAlbumURLs(string(body), pageURL)
}

func (this FilesterService) GetFileInfo(pageURL string) (string, string, error) {
    body, err := this.Client.Get(pageURL)
    if err != nil {
        return ``, ``, fmt.Errorf(`Failed to fetch page: %w`, err)
    }

    filename, err := this.ParseFileInfo(string(body))
    if err != nil {
        return "", "", fmt.Errorf("Failed parsing file info: %w", err)
    }
    
    parsed, err := url.Parse(pageURL)
    if err != nil {
        return ``, ``, fmt.Errorf(`Failed to parse URL: %w`, err)
    }
    origin := parsed.Scheme + `://` + parsed.Host
    slug := path.Base(parsed.Path)

    payload, _ := json.Marshal(map[string]string{"file_slug": slug})
    downloadURL, err := this.GetTokenData(pageURL, origin, payload)

    cdn := this.CDNURLs[rand.Intn(len(this.CDNURLs))]
    return filename, cdn + downloadURL + `?download=true`, nil
}

func (this FilesterService) GetTokenData(pageURL string, origin string, payload []byte) (string, error) {
    //
    resBody, err := this.Client.Post(pageURL, origin + "/api/public/download", payload)
    if err != nil {
        return "", fmt.Errorf("Token request failed: %w", err)
    }

    tokenData := Models.FilesterTokenData{}
    err = json.Unmarshal(resBody, &tokenData);
    if err != nil {
        return "", fmt.Errorf("Failed to parse token response: %w", err)
    }
    if tokenData.DownloadURL == "" {
        return "", fmt.Errorf("Empty download URL in response")
    }

    return tokenData.DownloadURL, nil
}
