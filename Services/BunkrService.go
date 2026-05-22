package Services

import (
	"regexp"

	"GDownloader/Interfaces"
	"GDownloader/Utils"
)

type BunkrService struct{
	//
	Name string
	SupportsRegex *regexp.Regexp
}

func (this BunkrService) Build() Interfaces.IDownloadProvider{
	//
	r, _ := regexp.Compile("")

	return BunkrService{ 
		SupportsRegex: r,
		Name: "BunkrService",
	}
}

func (this BunkrService) Download(url string){
	//
	Utils.Logger.Log("Downloading from " + this.Name)
}

func (this BunkrService) Supports(url string) (bool, error){
	//
	return regexp.MatchString(this.SupportsRegex.String(), url);
}
