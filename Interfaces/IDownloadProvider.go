package Interfaces

import "GDownloader/Utils"

type IDownloadProvider interface {
	//
	Build() IDownloadProvider
	Supports(url string) bool
	Download(client Utils.HTTPClient, filename string, pageURL string, downloadURL string, rootURL string) error
	HandleDownload(url string)
}
