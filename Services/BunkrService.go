package Services

import (
	"regexp"

	"GDownloader/Interfaces"
	"GDownloader/Utils"
)

type BunkrService struct{
	//
	Name string
	BaseURL string
	CDNURLs []string
	SupportsRegex *regexp.Regexp
	CurrentInstances uint8
}

func (this BunkrService) Build() Interfaces.IDownloadProvider{
	//
	return BunkrService{ 
		Name: "BunkrService", // bunkr.cr | bunkr.black | bunkr.site | bunkr.pk
		BaseURL: "bunkr.",
		CDNURLs: []string { 
			"https://get.bunkrr.su",
		},
		SupportsRegex: regexp.MustCompile(`(?i)^https?://bunkr\.[a-z]+`),
		CurrentInstances: 0,
	}
}

func (this BunkrService) Supports(url string) bool{
	//
	return this.SupportsRegex.MatchString(url)
}

func (this BunkrService) Download(url string) error{
	//
	Utils.Logger.Log("Downloading from " + this.Name)

	return nil
}

func (this BunkrService) HandleDownload(url string){
	//
	Utils.Logger.Log(this.Name + " " + url)
}
