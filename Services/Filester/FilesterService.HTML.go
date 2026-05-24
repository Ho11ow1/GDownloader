package Services

import (
	"strconv"

	"GDownloader/Models"
)

func (this FilesterService) GetSlugsFromPage(body string, origin string) []Models.FilesterSlug {
    //
    slugs := []Models.FilesterSlug{}
    
    matches := this.AlbumFileSlugRegex.FindAllStringSubmatch(body, -1)
    for _, match := range matches {
        if len(match) > 1 {
            slugs = append(slugs, Models.FilesterSlug{
                URL: origin + match[1],
                Filename: match[2],
            })
        }
    }

    return slugs
}

func (this FilesterService) GetPageCount(body string) int {
    //
    matches := this.PageCountRegex.FindStringSubmatch(body)
    if len(matches) < 2 {
        return 1
    }

    count, err := strconv.Atoi(matches[1])
    if err != nil {
        return 1
    }

    return count
}
