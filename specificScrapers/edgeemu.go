package specificScrapers

import (
	"hatt/helpers"
	"hatt/variables"
	"strings"

	"github.com/gocolly/colly"
)

func (t T) Edgeemu() []variables.Item {

	var results []variables.Item
	c := colly.NewCollector()

	config := helpers.DeserializeWebsiteConf("edgeemu.json")
	specificInfo := config.SpecificInfo
	itemKeys := config.Search.ItemKeys

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {

		item := variables.Item{
			Name: h.ChildText(itemKeys.Name),
			Link: h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
		}

		if item.Name != "" {
			item.Metadata = map[string]string{
				"console": h.ChildText(specificInfo["console"]),
				"size":    h.ChildText(specificInfo["size"]),
			}
			results = append(results, item)
		}
	})

	c.Visit(config.Search.Url + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement))

	return results
}
