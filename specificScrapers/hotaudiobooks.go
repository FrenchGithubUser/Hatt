package specificScrapers

import (
	"hatt/assets"
	"hatt/variables"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

func (t T) Hotaudiobooks() []variables.Item {

	var results []variables.Item
	c := colly.NewCollector()

	config := assets.DeserializeWebsiteConf("hotaudiobooks.json")
	itemKeys := config.Search.ItemKeys

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {

		item := variables.Item{
			Name: strings.ReplaceAll(h.ChildText(itemKeys.Name), ".", " "),
			Link: h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
		}

		results = append(results, item)
	})

	c.Visit(config.Search.Url + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement))

	var wg sync.WaitGroup
	for index, item := range results {

		wg.Add(1)
		go func(item variables.Item, index int) {
			c := colly.NewCollector()
			c.OnHTML("figure", func(h *colly.HTMLElement) {
				results[index].Thumbnail = h.ChildAttr(itemKeys.Thumbnail.Key, itemKeys.Thumbnail.Attribute)
			})
			c.Visit(item.Link)
			wg.Done()
		}(item, index)

	}
	wg.Wait()

	return results
}
