package Models

import (
	"regexp"

	"GDownloader/Common"
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
