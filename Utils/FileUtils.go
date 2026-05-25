package Utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"GDownloader/Common"
)

func GetAvailableDestinationPath(filename string, pageURL string) string {
	//
	subDir := filepath.Join(Common.AppDefs.DownloadDir, filepath.Base(pageURL))
	os.MkdirAll(subDir, os.ModePerm)
    destPath := filepath.Join(subDir, filepath.Base(filename))
    
    _, err := os.Stat(destPath); 
	if err == nil {
		ext := filepath.Ext(filename)
		name := strings.TrimSuffix(filename, ext)

		n := 1
		for {
			newFilename := fmt.Sprintf(`%s (%d)%s`, name, n, ext)
			destPath = filepath.Join(subDir, newFilename)

			_, err := os.Stat(destPath);
			if err != nil {
				break
			}

			n++
		}
	}

	return destPath
}
