package Services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"GDownloader/Common"
	"GDownloader/Models"
	"GDownloader/Utils"
)

func (this BunkrService) Download(client Utils.HTTPClient, filename string, pageURL string, downloadURL string, rootURL string) error {
	//
	destPath := Utils.GetAvailableDestinationPath(filename, rootURL)

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

func (this BunkrService) HandleDownload(pageURL string) {
	//
    client := Utils.HTTPClient{Client: &http.Client{}}

    if this.Base.IsAlbum(Common.Bunkr, pageURL) {
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

            go func(s Models.BunkrSlug) {
                defer wg.Done()
                defer func() { <-semaphore }()

                _, downloadURL, err := this.GetFileInfo(client, s.URL)
                if err != nil {
                    Utils.Logger.LogError(fmt.Sprintf("Failed to get download URL: %s", err))
                    return
                }

				Utils.Logger.Log(fmt.Sprintf("%s | %s | %s\n", s.Filename, s.URL, downloadURL))


				err = this.Download(client, s.Filename, s.URL, downloadURL, pageURL);
                if err != nil {
                    Utils.Logger.LogError(fmt.Sprintf("Download failed: %s", err))
                }
            }(slug)
        }

        wg.Wait()

    } else {
        filename, downloadURL, err := this.GetFileInfo(client, pageURL)
        if err != nil {
            Utils.Logger.LogError(fmt.Sprintf("Failed to get download URL: %s", err))
            return
        }

		Utils.Logger.Log(fmt.Sprintf("%s | %s", filename, downloadURL))
        
		err = this.Download(client, filename, pageURL, downloadURL, pageURL);
        if err != nil {
            Utils.Logger.LogError(fmt.Sprintf("Download failed: %s", err))
        }
    }
}

func (this BunkrService) GetAlbumURLs(client Utils.HTTPClient, pageURL string) ([]Models.BunkrSlug, error) {
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

func (this BunkrService) GetFileInfo(client Utils.HTTPClient, pageURL string) (string, string, error) {
	//
    body, err := client.Get(pageURL)
    if err != nil {
        return "", "", fmt.Errorf("Failed to fetch page: %w", err)
    }
    bodyStr := string(body)

    nameMatches := this.SingleFileNameRegex.FindStringSubmatch(bodyStr)
    if len(nameMatches) < 2 {
        return "", "", fmt.Errorf("Could not find filename on page")
    }

    idMatches := this.DownloadIDRegex.FindStringSubmatch(bodyStr)
    if len(idMatches) < 2 {
        return "", "", fmt.Errorf("Could not find download ID on page")
    }

    payload, _ := json.Marshal(map[string]string{"id": idMatches[1]})
    resBody, err := client.Post(pageURL, "https://apidl.bunkr.ru/api/_001_v2", payload)
    if err != nil {
        return "", "", fmt.Errorf("API request failed: %w", err)
    }

    tokenData := Models.BunkrTokenData{}
    if err = json.Unmarshal(resBody, &tokenData); err != nil {
        return "", "", fmt.Errorf("Failed to parse token response: %w", err)
    }

    downloadURL, err := this.DecryptURL(tokenData.URL, tokenData.Timestamp)
    if err != nil {
        return "", "", err
    }

    return nameMatches[1], downloadURL + "?n=" + nameMatches[1], nil
}
