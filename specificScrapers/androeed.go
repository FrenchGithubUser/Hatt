package specificScrapers

import (
	"hatt/helpers"
	"hatt/variables"
	"strings"

	"github.com/gocolly/colly"
)

func (t T) Androeed() []variables.Item {

	var results []variables.Item
	c := colly.NewCollector()

	config := helpers.DeserializeWebsiteConf("androeed.json")
	itemKeys := config.Search.ItemKeys

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {

		item := variables.Item{
			Name:      h.ChildText(itemKeys.Name),
			Thumbnail: h.ChildAttr(itemKeys.Thumbnail.Key, itemKeys.Thumbnail.Attribute),
			Link:      h.Request.AbsoluteURL(h.Attr("href")),
		}

		if item.Name != "" {

			item.Metadata = map[string]string{
				"size": h.ChildText(".size"),
			}

			results = append(results, item)
		}
	})

	c.Visit(config.Search.Url + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement))

	return results
}
