package Services

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"GDownloader/Models"
)

func (this BunkrService) ParseSlugsFromPage(body string, origin string) []Models.BunkrSlug {
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

func (this BunkrService) ParsePageCount(body string) int {
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

func (this BunkrService) ParseAlbumURLs(body string, pageURL string) ([]Models.BunkrSlug, error) {
    //
    parsed, err := url.Parse(pageURL)
    if err != nil {
        return nil, err
    }
    origin := parsed.Scheme + "://" + parsed.Host

    allSlugs := this.ParseSlugsFromPage(string(body), origin)
    pageCount := this.ParsePageCount(string(body))
    for page := 2; page <= pageCount; page++ {
        pageBody, err := this.Client.Get(strings.Split(pageURL, "?")[0] + fmt.Sprintf("?page=%d", page))
        if err != nil {
            return nil, err
        }

        allSlugs = append(allSlugs, this.ParseSlugsFromPage(string(pageBody), origin)...)
    }

    return allSlugs, nil
}

func (this BunkrService) ParseFileInfo(body string) (string, string, error) {
    //
    nameMatches := this.SingleFileNameRegex.FindStringSubmatch(body)
    if len(nameMatches) < 2 {
        return "", "", fmt.Errorf("Could not find filename on page")
    }

    idMatches := this.DownloadIDRegex.FindStringSubmatch(body)
    if len(idMatches) < 2 {
        return "", "", fmt.Errorf("Could not find download ID on page")
    }

    return nameMatches[1], idMatches[1], nil
}
