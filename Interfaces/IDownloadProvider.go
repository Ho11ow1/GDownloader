package Interfaces

type IDownloadProvider interface {
	//
	Build() IDownloadProvider
	Supports(url string) bool
	Download(filename string, pageURL string, downloadURL string, rootURL string) error
	HandleDownload(url string)
}
