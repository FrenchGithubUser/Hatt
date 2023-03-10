package specificScrapers

import (
	"hatt/assets"
	"hatt/variables"
	"strings"

	"github.com/gocolly/colly"
)

func (t T) Libgenrs() []variables.Item {

	var results []variables.Item
	c := colly.NewCollector()

	config := assets.DeserializeWebsiteConf("libgenrs.json")
	itemKeys := config.Search.ItemKeys

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {

		item := variables.Item{
			Name:      h.ChildText(itemKeys.Name),
			Thumbnail: config.SpecificInfo["baseUrl"] + h.ChildAttr(itemKeys.Thumbnail.Key, itemKeys.Thumbnail.Attribute),
			Link:      h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
		}

		if item.Name != "" {

			item.Metadata = map[string]string{
				"size":     h.ChildText(config.SpecificInfo["size"]),
				"language": h.ChildText(config.SpecificInfo["language"]),
			}

			year := h.ChildText(config.SpecificInfo["year"])
			if year != "" {
				item.Metadata["year"] = year
			}

			results = append(results, item)
		}
	})

	c.Visit(config.Search.Url + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement))

	return results
}
