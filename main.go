package main

import (
	"fmt"
	"strings"

	"GDownloader/Common"
	"GDownloader/Services"
	"GDownloader/Utils"
)

func main() {
	//
	var str strings.Builder
	for _, url := range Common.AppConfig.Urls {
		str.WriteString(url)
	}

	Dispatch()

	fmt.Printf("%05d | %s | %s | %s", Utils.PtrTernary(Common.AppConfig.Limit, 00000), Utils.PtrTernary(Common.AppConfig.Prefix, "None"), Utils.PtrTernary(Common.AppConfig.Extension, "None"), str.String())
}

func Dispatch() {
	//
	urls := Common.AppConfig.Urls

	for _, url := range urls {
		switch Utils.GetBaseUrl(url) {
			//
			case Common.AppDefs.BunkrBaseURL:
				Services.BunkrService.Build(Services.BunkrService{}).Download(url)

			case Common.AppDefs.FilesterBaseURL:
				Services.FilesterService.Build(Services.FilesterService{}).Download(url)
				
			case Common.AppDefs.FileDitchBaseURL:
				Services.FileDitchService.Build(Services.FileDitchService{}).Download(url)
		}
	}
}
