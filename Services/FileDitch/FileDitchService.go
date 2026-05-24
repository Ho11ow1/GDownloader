package Services

import (
	"regexp"

	"GDownloader/Interfaces"
	"GDownloader/Models"
)

type FileDitchService struct {
	//
	Base Models.DownloadServiceBase
	Name string
	BaseURL string
	CDNURLs []string
	SupportsRegex *regexp.Regexp
	DownloadURLRegex *regexp.Regexp
}

func (this FileDitchService) Build() Interfaces.IDownloadProvider {
	//
	return FileDitchService{ 
		Base: Models.DownloadServiceBase{},
		Name: "FileDitchService",
		BaseURL: "fileditchfiles.me",
		CDNURLs: []string {
			"https://1.thegumonmyshoe.me",
			"https://thegumonmyshoe.me",
		},
		SupportsRegex: regexp.MustCompile(`(?i)^https?://(www\.)?fileditchfiles\.me`),
		DownloadURLRegex: regexp.MustCompile(`<a\s+href="([^"]+)"[^>]+class="btn btn-main"[^>]+id="d[^"]*"[^>]+download`),
	}
}

func (this FileDitchService) Supports(url string) bool {
	//
	return this.SupportsRegex.MatchString(url)
}
