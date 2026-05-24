package Services

import (
	"GDownloader/Interfaces"

	bunkr "GDownloader/Services/Bunkr"
	fileditch "GDownloader/Services/FileDitch"
	filester "GDownloader/Services/Filester"
)

var AvailableServices = []Interfaces.IDownloadProvider {
	//
	bunkr.BunkrService{}.Build(),
	fileditch.FileDitchService{}.Build(),
	filester.FilesterService{}.Build(),
}
