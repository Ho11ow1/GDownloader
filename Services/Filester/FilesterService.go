package Services

import (
	"regexp"

	"GDownloader/Interfaces"
	"GDownloader/Models"
)

type FilesterService struct {
	//
	Base Models.DownloadServiceBase
	Name string
	BaseURL string
	CDNURLs []string
	SupportsRegex *regexp.Regexp
	AlbumFileSlugRegex *regexp.Regexp
	PageCountRegex *regexp.Regexp
	SingleFileNameRegex *regexp.Regexp
}

func (this FilesterService) Build() Interfaces.IDownloadProvider {
	//
	return FilesterService{
		Base: Models.DownloadServiceBase{},
		Name: "FilesterService",
		BaseURL: "filester.", // filester.si | filester.gg | filester.me | filester.sh
		CDNURLs: []string {
			"https://cache1.filester.me",
			"https://cache6.filester.me",
			"https://cache00.filester.me",
			"https://cn1.filester.me",
		},
		SupportsRegex: regexp.MustCompile(`(?i)^https?://filester\.[a-z]+`),
		AlbumFileSlugRegex: regexp.MustCompile(`data-name="([^"]+)"[^>]+onclick="window\.location\.href='(/d/[^']+)'"`),
		PageCountRegex: regexp.MustCompile(`<span class="page-info">\d+ / (\d+)</span>`),
		SingleFileNameRegex: regexp.MustCompile(`window\.fileName\s*=\s*"([^"]+)"`),
	}
}

func (this FilesterService) Supports(url string) bool {
	//
	return this.SupportsRegex.MatchString(url)
}
