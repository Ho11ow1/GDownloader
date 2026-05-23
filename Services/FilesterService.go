package Services

import (
	"regexp"

	"GDownloader/Common"
	"GDownloader/Interfaces"
	"GDownloader/Utils"
)

type FilesterService struct{
	//
	Name string
	BaseURL string
	CDNURLs []string
	SupportsRegex *regexp.Regexp
	CurrentInstances uint8
}

func (this FilesterService) Build() Interfaces.IDownloadProvider{
	//
	return FilesterService{ 
		Name: "FilesterService",
		BaseURL: "filester.", // filester.si | filester.gg | filester.me | filester.sh
		CDNURLs: []string {
			"https://cache1.filester.me",
			"https://cache6.filester.me",
			"https://cache00.filester.me",
			"https://cn1.filester.me",
		},
		SupportsRegex: regexp.MustCompile(`(?i)^https?://filester\.[a-z]+`),
		CurrentInstances: 0,
	}
}

func (this FilesterService) Supports(url string) bool{
	//
	return this.SupportsRegex.MatchString(url)
}

func (this FilesterService) Download(url string) error{
	//
	body, err := Utils.HTTPClient{}.Get(url, Common.AppDefs.Timeout)
	if err != nil {
		panic("Ran into http client get error")
	}

	Utils.Logger.LogSuccess(string(body))

	return nil
}

func (this FilesterService) HandleDownload(url string){
	//
	Utils.Logger.Log(this.Name + " " + url)
	Utils.Logger.Log("")
	
	this.Download(url)
}
