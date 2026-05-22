package Utils

import (
	"net/url"
	"strings"

	"GDownloader/Common"
)

func GetBaseUrl(rawURL string) string{
	//
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}

	host := u.Hostname()

	switch {
		case strings.HasPrefix(host, Common.AppDefs.BunkrBaseURL):
			return Common.AppDefs.BunkrBaseURL

		case strings.HasPrefix(host, Common.AppDefs.FilesterBaseURL):
			return Common.AppDefs.FilesterBaseURL

		case host == Common.AppDefs.FileDitchBaseURL:
			return Common.AppDefs.FileDitchBaseURL
	}

	return ""
}
