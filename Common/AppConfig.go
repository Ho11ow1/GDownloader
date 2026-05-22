package Common

type appConfig struct{
	//
	Urls []string
	Limit *uint
	Prefix *string
	Extension *string
}

var AppConfig = &appConfig{}
