package Services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"path"
	"strings"
	"sync"
	"time"

	"GDownloader/Common"
	"GDownloader/Models"
	"GDownloader/Utils"
)

func (this FilesterService) Download(client Utils.HTTPClient, filename string, pageURL string, downloadURL string) error {
	//
	destPath := Utils.GetAvailableDestinationPath(filename)

    var err error
    for attempt := 1; attempt <= int(Common.AppDefs.MaxRetry); attempt++ {
        err = client.Download(pageURL, downloadURL, destPath)
        if err == nil {
            Utils.Logger.LogSuccess(fmt.Sprintf("Successfully downloaded: %s", filename))
            return nil
        }

        Utils.Logger.LogError(fmt.Sprintf("Download attempt %d failed: %s", attempt, err))
        time.Sleep(Common.AppDefs.RetryDelay * time.Duration(attempt))
    }

    return fmt.Errorf("All attempts failed: %w", err)
}

func (this FilesterService) HandleDownload(pageURL string) {
    //
    client := Utils.HTTPClient{ Client: &http.Client{}}

    if this.Base.IsAlbum(Common.Filester, pageURL) {
        urls, err := this.GetAlbumURLs(client, pageURL)
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

                _, downloadURL, err := this.GetFileInfo(client, s.URL)
                if err != nil {
                    Utils.Logger.LogError(fmt.Sprintf("Failed to get download URL: %s", err))
                    return
                }

                err = this.Download(client, s.Filename, s.URL, downloadURL)
                if err != nil {
                    Utils.Logger.LogError(fmt.Sprintf("Download failed: %s", err))
                }
            }(slug)
        }

        wg.Wait()

    } else {
        filename, downloadURL, err := this.GetFileInfo(client, pageURL)
        if err != nil {
            Utils.Logger.LogError(fmt.Sprintf("Failed to get file info: %s", err))
            return
        }

        err = this.Download(client, filename, pageURL, downloadURL);
        if err != nil {
            Utils.Logger.LogError(fmt.Sprintf("Download failed: %s", err))
        }
    }
}

func (this FilesterService) GetAlbumURLs(client Utils.HTTPClient, pageURL string) ([]Models.FilesterSlug, error) {
    //
    origin, err := Utils.ParseOrigin(pageURL)
    if err != nil {
        return nil, err
    }

    body, err := client.Get(strings.Split(pageURL, "?")[0] + "?page=1")
    if err != nil {
        return nil, fmt.Errorf("Failed to fetch album page 1: %w", err)
    }

    allSlugs := this.GetSlugsFromPage(string(body), origin)
    pageCount := this.GetPageCount(string(body))
    for page := 2; page <= pageCount; page++ {
        pageBody, err := client.Get(strings.Split(pageURL, "?")[0] + fmt.Sprintf("?page=%d", page))
        if err != nil {
            return nil, fmt.Errorf("Failed to fetch album page %d: %w", page, err)
        }

        allSlugs = append(allSlugs, this.GetSlugsFromPage(string(pageBody), origin)...)
    }

    return allSlugs, nil
}

func (this FilesterService) GetFileInfo(client Utils.HTTPClient, pageURL string) (string, string, error) {
    //
    origin, err := Utils.ParseOrigin(pageURL)
    if err != nil {
        return "", "", err
    }

    body, err := client.Get(pageURL)
    if err != nil {
        return "", "", fmt.Errorf("Failed to fetch page: %w", err)
    }

    bodyStr := string(body)
    filenameMatches := this.SingleFileNameRegex.FindStringSubmatch(bodyStr)
    if len(filenameMatches) < 2 {
        return "", "", fmt.Errorf("Could not find filename on page")
    }

    slug := path.Base(filenameMatches[0])
    payload, _ := json.Marshal(map[string]string{"file_slug": slug})
    resBody, err := client.Post(pageURL, origin + "/api/public/download", payload)
    if err != nil {
        return "", "", fmt.Errorf("Token request failed: %w", err)
    }

    tokenData := Models.FilesterTokenData{}
    err = json.Unmarshal(resBody, &tokenData); 
    if err != nil{
        return "", "", fmt.Errorf("Failed to parse token response: %w", err)
    }
    if tokenData.DownloadURL == "" {
        return "", "", fmt.Errorf("Empty download URL in response")
    }

    cdn := this.CDNURLs[rand.Intn(len(this.CDNURLs))]
    return filenameMatches[1], cdn + tokenData.DownloadURL + "?download=true", nil
}
