package main

import (
	"fmt"
	"hatt/configuration"
	"strings"

	"github.com/gocolly/colly"
)

func scrapePlainHtml(input string, config configuration.Config) []item {

	var items []item
	c := colly.NewCollector()
	itemKeys := config.Search.ItemKeys

	// c.OnHTML("body", func(h *colly.HTMLElement) {
	// 	fmt.Println(h)
	// })

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {

		item := item{
			Name:      h.ChildText(itemKeys.Name),
			Thumbnail: h.ChildAttr(itemKeys.Thumbnail.Key, itemKeys.Thumbnail.Attribute),
			Link:      h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
		}

		items = append(items, item)
	})

	// pagination handling
	// c.OnHTML("a.navigation.next", func(h *colly.HTMLElement) {
	// 	nextPage := h.Request.AbsoluteURL(h.Attr("href"))
	// 	c.Visit(nextPage)

	// })

	// when website requires login
	if config.Login.Url != "" {
		// login(config.Name)
		// tokens := deserializeCredentials()[config.Name]["tokens"]
		c.OnRequest(func(r *colly.Request) {
			// for _, token := range config.Login.Tokens {
			// }
			r.Headers.Set("cookie", "")
		})
	}

	fmt.Println(config.Search.Url + strings.ReplaceAll(input, " ", config.Search.SpaceReplacement))
	c.Visit(config.Search.Url + strings.ReplaceAll(input, " ", config.Search.SpaceReplacement))

	return items
}
