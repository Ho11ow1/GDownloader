package Services

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"GDownloader/Common"
	"GDownloader/Models"
	"GDownloader/Utils"
)

func (this BunkrService) HandleDownload(pageURL string) {
	//
    if this.Base.IsAlbum(Common.Bunkr, pageURL) {
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

            go func(s Models.BunkrSlug) {
                defer wg.Done()
                defer func() { <-semaphore }()

                _, downloadURL, err := this.GetFileInfo(s.URL)
                if err != nil {
                    Utils.Logger.LogError(fmt.Sprintf("Failed to get download URL: %s", err))
                    return
                }

				Utils.Logger.Log(fmt.Sprintf("%s | %s | %s\n", s.Filename, s.URL, downloadURL))


				err = this.Base.Download(this.Client, s.Filename, s.URL, downloadURL, pageURL);
                if err != nil {
                    Utils.Logger.LogError(fmt.Sprintf("Download failed: %s", err))
                }
            }(slug)
        }

        wg.Wait()

    } else {
        filename, downloadURL, err := this.GetFileInfo(pageURL)
        if err != nil {
            Utils.Logger.LogError(fmt.Sprintf("Failed to get download URL: %s", err))
            return
        }

		Utils.Logger.Log(fmt.Sprintf("%s | %s", filename, downloadURL))
        
		err = this.Base.Download(this.Client, filename, pageURL, downloadURL, pageURL);
        if err != nil {
            Utils.Logger.LogError(fmt.Sprintf("Download failed: %s", err))
        }
    }
}

func (this BunkrService) GetAlbumURLs(pageURL string) ([]Models.BunkrSlug, error) {
	//
	body, err := this.Client.Get(strings.Split(pageURL, "?")[0] + "?page=1")
    if err != nil {
        return nil, fmt.Errorf("Failed to fetch album page 1: %w", err)
    }

	return this.ParseAlbumURLs(string(body), pageURL)
}

func (this BunkrService) GetFileInfo(pageURL string) (string, string, error) {
	//
    body, err := this.Client.Get(pageURL)
    if err != nil {
        return "", "", fmt.Errorf("Failed to fetch page: %w", err)
    }

    name, id, err := this.ParseFileInfo(string(body))

    payload, _ := json.Marshal(map[string]string{"id": id})
    tokenURL, timestamp, err := this.GetTokenData(pageURL, payload)
    if err != nil {
        return "", "", err
    }

    downloadURL, err := this.DecryptURL(tokenURL, timestamp)
    if err != nil {
        return "", "", err
    }

    return name, downloadURL + "?n=" + url.PathEscape(name), nil
}

func (this BunkrService) GetTokenData(pageURL string, payload []byte) (string, int64, error){
    //
    resBody, err := this.Client.Post(pageURL, "https://apidl.bunkr.ru/api/_001_v2", payload)
    if err != nil {
        return "", 0, fmt.Errorf("API request failed: %w", err)
    }

    tokenData := Models.BunkrTokenData{}
    err = json.Unmarshal(resBody, &tokenData);
    if err != nil {
        return "", 0, fmt.Errorf("Failed to parse token response: %w", err)
    }

    return tokenData.URL, tokenData.Timestamp, nil
}
