package specificScrapers

import (
	"hatt/assets"
	"hatt/htmlParsers"
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
			Name: strings.ReplaceAll(h.ChildText(itemKeys.Name), ".", " "),
			Link: h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
		}
		item.Thumbnail = htmlParsers.ScrapeItemPageHtml(config, item.Link)

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
