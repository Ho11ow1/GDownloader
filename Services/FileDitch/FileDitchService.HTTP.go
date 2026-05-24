package Services

import (
	"fmt"
	"html"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"GDownloader/Common"
	"GDownloader/Utils"
)

func (this FileDitchService) Download(client Utils.HTTPClient, filename string, pageURL string, downloadURL string) error {
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

func (this FileDitchService) HandleDownload(pageURL string) {
	//
	client := Utils.HTTPClient{ Client: &http.Client{}}

	downloadURL, err := this.GetDownloadURL(client, pageURL)
	if err != nil {
        Utils.Logger.LogError(fmt.Sprintf("Failed to get download URL: %s", err))
        return
    }

	filename := strings.TrimSuffix(filepath.Base(pageURL), filepath.Ext(pageURL)) + filepath.Ext(pageURL)
	err = this.Download(client, filename, pageURL, downloadURL);
    if err != nil {
        Utils.Logger.LogError(fmt.Sprintf("Download failed: %s", err))
    }
}

func (this FileDitchService) GetDownloadURL(client Utils.HTTPClient, pageURL string) (string, error) {
	//
    body, err := client.Get(pageURL)
    if err != nil {
        return "", fmt.Errorf("Failed to fetch page: %w", err)
    }

    matches := this.DownloadURLRegex.FindStringSubmatch(string(body))
    if len(matches) < 2 {
        return "", fmt.Errorf("Could not find download link on page")
    }

    return html.UnescapeString(matches[1]), nil
}
