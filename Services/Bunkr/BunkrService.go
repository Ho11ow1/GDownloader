package Services

import (
	"encoding/base64"
	"fmt"
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
	AlbumFileSlugRegex *regexp.Regexp
	PageCountRegex *regexp.Regexp
	PageNumRegex *regexp.Regexp
	DownloadLinkRegex *regexp.Regexp
	DownloadIDRegex *regexp.Regexp
	SingleFileNameRegex *regexp.Regexp
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
		AlbumFileSlugRegex: regexp.MustCompile(`<p class="truncate theName[^"]*">([^<]+)</p>[\s\S]*?<a class="after:absolute[^"]*"\s+href="(/f/[a-zA-Z0-9]+)"`),
		PageCountRegex: regexp.MustCompile(`(?s)<nav class="pagination">(.*?)</nav>`),
		PageNumRegex: regexp.MustCompile(`\?page=(\d+)`),
		DownloadLinkRegex: regexp.MustCompile(`<a class="btn btn-main[^"]*"\s+href="([^"]+)"`),
		DownloadIDRegex: regexp.MustCompile(`href="https://get\.bunkrr\.su/file/(\d+)"`),
		SingleFileNameRegex: regexp.MustCompile(`<h1 class="[^"]*truncate[^"]*">([^<]+)</h1>`),
	}
}

func (this BunkrService) Supports(url string) bool {
	//
	return this.SupportsRegex.MatchString(url)
}

func (this BunkrService) DecryptURL(encryptedB64 string, timestamp int64) (string, error) {
	//
    data, err := base64.StdEncoding.DecodeString(encryptedB64)
    if err != nil {
        return "", fmt.Errorf("Failed to base64 decode: %w", err)
    }

    key := []byte(fmt.Sprintf("SECRET_KEY_%d", timestamp / 3600))
    result := make([]byte, len(data))
    for i := range data {
        result[i] = data[i] ^ key[i % len(key)]
    }

    return string(result), nil
}
