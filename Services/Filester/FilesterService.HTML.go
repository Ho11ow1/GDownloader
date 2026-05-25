package Services

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"GDownloader/Models"
)

func (this FilesterService) ParseSlugsFromPage(body string, origin string) []Models.FilesterSlug {
    //
    slugs := []Models.FilesterSlug{}
    
    matches := this.AlbumFileSlugRegex.FindAllStringSubmatch(body, -1)
    for _, match := range matches {
        if len(match) > 1 {
            slugs = append(slugs, Models.FilesterSlug{
                URL: origin + match[2],
                Filename: match[1],
            })
        }
    }

    return slugs
}

func (this FilesterService) ParsePageCount(body string) int {
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

func (this FilesterService) ParseAlbumURLs(body string, pageURL string) ([]Models.FilesterSlug, error) {
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
            return nil, fmt.Errorf("Failed to fetch album page %d: %w", page, err)
        }

        allSlugs = append(allSlugs, this.ParseSlugsFromPage(string(pageBody), origin)...)
    }

    return allSlugs, nil
}

func (this FilesterService) ParseFileInfo(body string) (string, error) {
    //
    filenameMatches := this.SingleFileNameRegex.FindStringSubmatch(body)
    if len(filenameMatches) < 2 {
        return "", fmt.Errorf("Could not find filename on page")
    }

    return filenameMatches[1], nil
}
