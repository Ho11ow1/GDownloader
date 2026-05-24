package Services

import (
	"regexp"

	"GDownloader/Interfaces"
	"GDownloader/Models"
)

type BunkrService struct {
	//
	Base Models.DownloadServiceBase
	Name string
	BaseURL string
	CDNURLs []string
	SupportsRegex *regexp.Regexp
}

func (this BunkrService) Build() Interfaces.IDownloadProvider {
	//
	return BunkrService{ 
		Base: Models.DownloadServiceBase{},
		Name: "BunkrService", // bunkr.cr | bunkr.black | bunkr.site | bunkr.pk
		BaseURL: "bunkr.",
		CDNURLs: []string { 
			"https://get.bunkrr.su",
		},
		SupportsRegex: regexp.MustCompile(`(?i)^https?://bunkr\.[a-z]+`),
	}
}

func (this BunkrService) Supports(url string) bool {
	//
	return this.SupportsRegex.MatchString(url)
}
