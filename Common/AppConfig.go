package Common

import "sync/atomic"

type appConfig struct {
	//
	Urls []string
	Limit *atomic.Uint64
	Prefix *string
	Extension *string
}

var AppConfig = &appConfig{}
