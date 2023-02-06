package htmlParsers

import (
	"hatt/configuration"
	"hatt/variables"
	"strings"

	"github.com/gocolly/colly"
)

func ScrapePlainHtml(config configuration.Config) []variables.Item {

	var items []variables.Item
	c := colly.NewCollector()
	itemKeys := config.Search.ItemKeys

	// c.OnHTML("body", func(h *colly.HTMLElement) {
	// 	fmt.Println(h)
	// })

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {

		item := variables.Item{
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
		// tokens := helpers.DeserializeCredentials()[config.Name]["tokens"]
		c.OnRequest(func(r *colly.Request) {
			// for _, token := range config.Login.Tokens {
			// }
			r.Headers.Set("cookie", "")
		})
	}

	c.Visit(config.Search.Url + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement))

	return items
}
