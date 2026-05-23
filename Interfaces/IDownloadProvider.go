package Interfaces

type IDownloadProvider interface{
	//
	Build() IDownloadProvider
	Supports(url string) bool
	// Per file download
	Download(url string) error
	// Handles concurrency dispatching & error handling
	HandleDownload(url string)
}
