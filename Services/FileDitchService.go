package Services

import (
	"regexp"

	"GDownloader/Interfaces"
	"GDownloader/Utils"
)

type FileDitchService struct{
	//
	Name string
	SupportsRegex *regexp.Regexp
}

func (this FileDitchService) Build() Interfaces.IDownloadProvider{
	//
	r, _ := regexp.Compile("")

	return FileDitchService{ 
		SupportsRegex: r,
		Name: "FileDitchService",
	}
}

func (this FileDitchService) Download(url string){
	//
	Utils.Logger.Log("Downloading from " + this.Name)
}

func (this FileDitchService) Supports(url string) (bool, error){
	//
	return regexp.MatchString(this.SupportsRegex.String(), url);
}
