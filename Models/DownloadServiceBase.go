package Models

import (
	"fmt"
	"html"
	"regexp"
	"strings"
	"time"

	"GDownloader/Common"
	"GDownloader/Utils"
)

type DownloadServiceBase struct {}

var albumPatterns = map[Common.ServiceType]AlbumPattern {
	//
    Common.Filester: {
        Album:  regexp.MustCompile("/f/"),
        Single: regexp.MustCompile("/d/"),
    },
    Common.Bunkr: {
        Album:  regexp.MustCompile("/a/"),
        Single: regexp.MustCompile("/f/"),
    },
}

func (this DownloadServiceBase) IsAlbum(serviceType Common.ServiceType, pageURL string) bool {
	//
	pattern, ok := albumPatterns[serviceType]
	if !ok {
		return false
	}

	return pattern.Album.MatchString(pageURL)
}

func (this DownloadServiceBase) Download(client Utils.HTTPClient, filename string, pageURL string, downloadURL string, rootURL string) error {
	//
    if Common.AppConfig.Prefix != nil {
        if !strings.HasPrefix(filename, *Common.AppConfig.Prefix) {
            return nil
        }
    }
    if Common.AppConfig.Extension != nil {
        if !strings.HasSuffix(filename, *Common.AppConfig.Extension) {
            return nil
        }
    }
    if Common.AppConfig.Limit != nil{
        current := Common.AppConfig.Limit.Load()
        if current == 0 {
            return nil
        }

        Common.AppConfig.Limit.CompareAndSwap(current, current -1)
    }

	destPath := Utils.GetAvailableDestinationPath(html.UnescapeString(filename), rootURL)

    var err error
    for attempt := 1; attempt <= int(Common.AppDefs.MaxRetry); attempt++ {
        err = client.Download(pageURL, downloadURL, destPath)
        if err == nil {
            Utils.Logger.LogSuccess(fmt.Sprintf("Successfully downloaded: %s", html.UnescapeString(filename)))
            return nil
        }

        Utils.Logger.LogError(fmt.Sprintf("Download attempt %d failed: %s", attempt, err))
        time.Sleep(Common.AppDefs.RetryDelay * time.Duration(attempt))
    }

    if Common.AppConfig.Limit != nil {
        Common.AppConfig.Limit.Add(1)
    }

    return fmt.Errorf("All attempts failed: %w", err)
}
