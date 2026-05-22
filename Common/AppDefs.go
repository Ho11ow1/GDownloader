package Common

type appDefs struct{
	BunkrBaseURL string
	BunkrCDNURLs []string
	FileDitchBaseURL string
	FileDitchCDNURLs []string
	FilesterBaseURL string
	FilesterCDNURLs []string
}

var AppDefs = &appDefs{
	BunkrBaseURL: "bunkr.", // bunkr.cr | bunkr.black | bunkr.site | bunkr.pk
	BunkrCDNURLs: []string { 
		"https://get.bunkrr.su",
	},
	
	FileDitchBaseURL: "fileditchfiles.me",
	FileDitchCDNURLs: []string {
		"https://1.thegumonmyshoe.me",
	},

	FilesterBaseURL: "filester.", // filester.si | filester.gg | filester.me | filester.sh
	FilesterCDNURLs: []string {
		"https://cache1.filester.me",
		"https://cache6.filester.me",
		"https://cache00.filester.me",
		"https://cn1.filester.me",
	},
}
