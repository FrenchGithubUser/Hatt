package specificScrapers

import (
	"hatt/assets"
	"hatt/variables"
	"strings"

	"github.com/gocolly/colly"
)

func (t T) Ddlbase() []variables.Item {

	var results []variables.Item
	c := colly.NewCollector()

	config := assets.DeserializeWebsiteConf("ddlbase.json")
	itemKeys := config.Search.ItemKeys

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {

		item := variables.Item{
			Name:      strings.ReplaceAll(h.ChildText(itemKeys.Name), ".", " "),
			Thumbnail: h.ChildAttr(itemKeys.Thumbnail.Key, itemKeys.Thumbnail.Attribute),
			Link:      h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
		}

		item.Metadata = map[string]string{
			"host":        h.ChildText(itemKeys.Metadata["host"]),
			"forum":       h.ChildText(itemKeys.Metadata["forum"]),
			"publication": h.ChildText(itemKeys.Metadata["publication"]),
		}
		item.Name = strings.Trim(item.Name, item.Metadata["host"])

		results = append(results, item)
	})

	c.Visit(config.Search.Url + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement))

	return results
}
