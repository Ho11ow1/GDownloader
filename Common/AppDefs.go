package Common

import "time"

type appDefs struct{
	MaxConcurrent uint8
	Timeout time.Duration
}

var AppDefs = &appDefs{
	MaxConcurrent: 5,
	Timeout: 5 * time.Second,
}
