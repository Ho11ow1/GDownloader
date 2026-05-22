package Interfaces

type IDownloadProvider interface{
	//
	Download(url string)
	Supports(url string) (bool, error)
	Build() IDownloadProvider
}
