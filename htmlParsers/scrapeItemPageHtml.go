package htmlParsers

import (
	"hatt/configuration"
	"regexp"
	"time"

	"github.com/gocolly/colly"
)

func ScrapeItemPageHtml(config configuration.Config, url string) string {
	var thumbnailLink string
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; rv:109.0) Gecko/20100101 Firefox/109.0"
	itemKeys := config.Search.ItemKeys

	// c.OnHTML("body", func(h *colly.HTMLElement) {
	// 	fmt.Println(h)
	// })

	c.OnHTML(itemKeys.Thumbnail.Key, func(h *colly.HTMLElement) {
		if thumbnailLink == "" {
			if itemKeys.Thumbnail.Attribute == "style" {
				urlRegex := regexp.MustCompile(`url\(([^)]+)\)`)
				style := h.Attr(itemKeys.Thumbnail.Attribute)
				thumbnailLink = urlRegex.FindStringSubmatch(style)[1]
			} else {
				thumbnailLink = h.Attr(itemKeys.Thumbnail.Attribute)
			}

			if itemKeys.Thumbnail.AppendToSiteUrl {
				thumbnailLink = h.Request.AbsoluteURL(thumbnailLink)
			}
		}
	})

	// c.OnError(func(r *colly.Response, err error) { fmt.Println(r.StatusCode, err, r.Request.URL) })
	c.SetRequestTimeout(30 * time.Second)
	c.OnRequest(func(r *colly.Request) {
		// maybe set this value according to the user's locale ?
		r.Headers.Set("Accept-Language", "en")
	})
	c.Visit(url)

	return thumbnailLink
}
