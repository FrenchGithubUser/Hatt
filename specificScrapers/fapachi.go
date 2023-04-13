package specificScrapers

import (
	"hatt/assets"
	"hatt/variables"
	"strings"

	"github.com/gocolly/colly"
)

// returns scrambled file
func (t T) Fapachi() []variables.Item {

	var results []variables.Item
	c := colly.NewCollector()

	config := assets.DeserializeWebsiteConf("fapachi.json")
	specificInfo := config.SpecificInfo
	itemKeys := config.Search.ItemKeys

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {

		item := variables.Item{
			Name:      h.ChildText(itemKeys.Name),
			Thumbnail: specificInfo["siteUrl"] + h.ChildAttr(itemKeys.Thumbnail.Key, itemKeys.Thumbnail.Attribute),
			Link:      h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
		}

		c.OnHTML(".col-12 p", func(h *colly.HTMLElement) {
			item.Metadata = map[string]string{
				"files": strings.Split(h.Text, "Media: ")[1] + " files",
			}
		})

		c.Visit(item.Link)

		results = append(results, item)
	})

	var payload map[string]string = map[string]string{"searchVal": variables.CURRENT_INPUT}
	c.Post(config.Search.Url, payload)

	return results
}
