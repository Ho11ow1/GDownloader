package Services

import (
	"fmt"
	"html"
	"path/filepath"
	"strings"

	"GDownloader/Utils"
)

func (this FileDitchService) HandleDownload(pageURL string) {
	//
	downloadURL, err := this.GetDownloadURL(pageURL)
	if err != nil {
        Utils.Logger.LogError(fmt.Sprintf("Failed to get download URL: %s", err))
        return
    }

	filename := strings.TrimSuffix(filepath.Base(pageURL), filepath.Ext(pageURL)) + filepath.Ext(pageURL)
	err = this.Base.Download(this.Client, filename, pageURL, downloadURL, pageURL);
    if err != nil {
        Utils.Logger.LogError(fmt.Sprintf("Download failed: %s", err))
    }
}

func (this FileDitchService) GetDownloadURL(pageURL string) (string, error) {
	//
    body, err := this.Client.Get(pageURL)
    if err != nil {
        return "", fmt.Errorf("Failed to fetch page: %w", err)
    }

    matches := this.DownloadURLRegex.FindStringSubmatch(string(body))
    if len(matches) < 2 {
        return "", fmt.Errorf("Could not find download link on page")
    }

    return html.UnescapeString(matches[1]), nil
}
