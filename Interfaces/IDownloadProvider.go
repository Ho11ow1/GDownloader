package Interfaces

type IDownloadProvider interface {
	//
	Build() IDownloadProvider
	Supports(url string) bool
	HandleDownload(url string)
}
