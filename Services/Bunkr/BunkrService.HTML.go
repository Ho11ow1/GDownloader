package Services

import (
	"strconv"

	"GDownloader/Models"
)

func (this BunkrService) GetSlugsFromPage(body string, origin string) []Models.BunkrSlug {
    //
    slugs := []Models.BunkrSlug{}

    matches := this.AlbumFileSlugRegex.FindAllStringSubmatch(body, -1)
    for _, match := range matches {
        if len(match) > 2 {
            slugs = append(slugs, Models.BunkrSlug{
                Filename: match[1],
                URL: origin + match[2],
            })
        }
    }
    return slugs
}

func (this BunkrService) GetPageCount(body string) int {
	//
    navMatch := this.PageCountRegex.FindStringSubmatch(body)
    if len(navMatch) < 2 {
        return 1
    }

    matches := this.PageNumRegex.FindAllStringSubmatch(navMatch[1], -1)
    var max int
    for _, match := range matches {
        num, err := strconv.Atoi(match[1])
        if err == nil && num > max {
            max = num
        }
    }
    return max
}
