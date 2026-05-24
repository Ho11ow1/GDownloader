package main

import (
	"sync"

	"GDownloader/Common"
	"GDownloader/Interfaces"
	"GDownloader/Services"
)

func main() {
	//
	if len(Common.AppConfig.Urls) > 0 {
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
