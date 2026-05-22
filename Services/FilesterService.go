package Services

import (
	"regexp"
	"time"

	"GDownloader/Interfaces"
	"GDownloader/Utils"
)

type FilesterService struct{
	//
	Name string
	SupportsRegex *regexp.Regexp
}

func (this FilesterService) Build() Interfaces.IDownloadProvider{
	//
	r, _ := regexp.Compile("")

	return FilesterService{ 
		SupportsRegex: r,
		Name: "FilesterService",
	}
}

func (this FilesterService) Download(url string){
	//
	body, err := Utils.HTTPClient{}.Get(url, 3 * time.Second)
	if err != nil {
		panic("Ran into http client get error")
	}

	Utils.Logger.LogSuccess(string(body))

	Utils.Logger.Log("Downloading from " + this.Name)
}

func (this FilesterService) Supports(url string) (bool, error){
	//
	return regexp.MatchString(this.SupportsRegex.String(), url);
}
