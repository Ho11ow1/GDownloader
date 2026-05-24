package Common

import "time"

type appDefs struct {
	//
	MaxConcurrent uint8
	MaxRetry uint8
	RetryDelay time.Duration
	DownloadDir string
}

var AppDefs = &appDefs{
	MaxConcurrent: 5,
	MaxRetry: 3,
	RetryDelay: 3 * time.Second,
	DownloadDir: `./Downloads`,
}
