package Utils

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"GDownloader/Common"
)

func GetAvailableDestinationPath(filename string) string {
	//
	destPath := filepath.Join(Common.AppDefs.DownloadDir, filename)
    
    _, err := os.Stat(destPath); 
	if err == nil {
		ext := filepath.Ext(filename)
		name := strings.TrimSuffix(filename, ext)

		n := 1
		for {
			newFilename := fmt.Sprintf(`%s (%d)%s`, name, n, ext)
			destPath = filepath.Join(Common.AppDefs.DownloadDir, newFilename)

			_, err := os.Stat(destPath);
			if err != nil {
				break
			}

			n++
		}
	}

	return destPath
}

func ParseOrigin(pageURL string) (string, error) {
	//
    parsed, err := url.Parse(pageURL)
    if err != nil {
        return "", fmt.Errorf("Failed to parse URL: %w", err)
    }
	
    return parsed.Scheme + "://" + parsed.Host, nil
}
