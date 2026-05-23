package main

import (
	"fmt"

	"GDownloader/Common"
	"GDownloader/Interfaces"
	"GDownloader/Services"
	"GDownloader/Utils"
)

func main() {
	//
	fmt.Printf("%05d | %s | %s | %s\n", 
		Utils.PtrTernary(Common.AppConfig.Limit, 00000), 
		Utils.PtrTernary(Common.AppConfig.Prefix, "None"), 
		Utils.PtrTernary(Common.AppConfig.Extension, "None"), 
		Utils.ArrayToString(Common.AppConfig.Urls),
	)

	Dispatch()
}

func Dispatch() {
	//
	services := []Interfaces.IDownloadProvider{
		Services.BunkrService{}.Build(),
		Services.FilesterService{}.Build(),
		Services.FileDitchService{}.Build(),
	}

	for _, url := range Common.AppConfig.Urls{
		for _, service := range services{
			if service.Supports(url){
				service.HandleDownload(url)
			}
		}
	}
}
