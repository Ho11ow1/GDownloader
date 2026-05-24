package Services

import (
	"fmt"
	"net/http"
	"path"

	"GDownloader/Utils"
)

func (this BunkrService) Download(client Utils.HTTPClient, filename string, pageURL string, downloadURL string) error {
	//
	Utils.Logger.Log("Downloading from " + this.Name)

	return nil
}

func (this BunkrService) HandleDownload(pageURL string) {
	//
	client := Utils.HTTPClient{ Client: &http.Client{}}

	downloadURL, err := this.GetDownloadURL(client, pageURL)
	if err != nil {
        Utils.Logger.LogError(fmt.Sprintf("Failed to get download URL: %s", err))
        return
    }
	
	err = this.Download(client, pageURL, path.Base(pageURL), downloadURL);
    if err != nil {
        Utils.Logger.LogError(fmt.Sprintf("Download failed: %s", err))
    }
}

func (this BunkrService) GetDownloadURL(client Utils.HTTPClient, pageURL string) (string, error) {
	//
	body, err := client.Get(pageURL)
    if err != nil {
        return "", fmt.Errorf("Failed to fetch page: %w", err)
    }

	return string(body), nil
}