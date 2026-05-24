package Models

import "regexp"

type AlbumPattern struct {
	//
    Album  *regexp.Regexp
    Single *regexp.Regexp
}
