package Services

import (
	"regexp"

	"GDownloader/Interfaces"
	"GDownloader/Utils"
)

type FileDitchService struct{
	//
	Name string
	BaseURL string
	CDNURLs []string
	SupportsRegex *regexp.Regexp
	CurrentInstances uint8
}

func (this FileDitchService) Build() Interfaces.IDownloadProvider{
	//
	return FileDitchService{ 
		Name: "FileDitchService",
		BaseURL: "fileditchfiles.me",
		CDNURLs: []string {
			"https://1.thegumonmyshoe.me",
		},
		SupportsRegex: regexp.MustCompile(`(?i)^https?://(www\.)?fileditchfiles\.me`),
		CurrentInstances: 0,
	}
}

func (this FileDitchService) Supports(url string) bool{
	//
	return this.SupportsRegex.MatchString(url)
}

func (this FileDitchService) Download(url string) error{
	//
	Utils.Logger.Log("Downloading from " + this.Name)

	return nil
}

func (this FileDitchService) HandleDownload(url string){
	//
	Utils.Logger.Log(this.Name + " " + url)
}
