package specificScrapers

import (
	"hatt/assets"
	"hatt/variables"
	"strings"

	"github.com/gocolly/colly"
)

func (t T) Fapello() []variables.Item {

	var results []variables.Item
	c := colly.NewCollector()

	config := assets.DeserializeWebsiteConf("fapello.json")
	itemKeys := config.Search.ItemKeys

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {

		item := variables.Item{
			Name:      h.ChildText(itemKeys.Name),
			Thumbnail: h.ChildAttr(itemKeys.Thumbnail.Key, itemKeys.Thumbnail.Attribute),
			Link:      h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
		}

		c.OnHTML(".divide-gray-300", func(h *colly.HTMLElement) {
			item.Metadata = map[string]string{
				"files": strings.Split(h.ChildText(".flex"), "Media")[0] + "files",
			}
		})

		c.Visit(item.Link)

		results = append(results, item)
	})

	c.Visit(config.Search.Url + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement))

	return results
}
