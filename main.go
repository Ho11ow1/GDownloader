package main

import (
	"fmt"
	"os"
	"sync"

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

	if len(Common.AppConfig.Urls) > 0 {
		err := os.MkdirAll(Common.AppDefs.DownloadDir, os.ModePerm); 
		if err != nil {
			Utils.Logger.LogError("Unable to create directory at " + Common.AppDefs.DownloadDir)
		}

		Dispatch()
	}
}

func Dispatch() {
	//
	var wg sync.WaitGroup

	for _, url := range Common.AppConfig.Urls {
		for _, service := range Services.AvailableServices {
			if service.Supports(url) {
				wg.Add(1)
				go func(d Interfaces.IDownloadProvider, u string) {
					defer wg.Done()
					d.HandleDownload(u)
				}(service, url)
			}
		}
	}

	wg.Wait()
}
