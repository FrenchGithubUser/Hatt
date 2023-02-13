package specificScrapers

import (
	"hatt/helpers"
	"hatt/variables"
	"strings"

	"github.com/gocolly/colly"
)

func (t T) Vimm() []variables.Item {

	var results []variables.Item
	c := colly.NewCollector()

	config := helpers.DeserializeWebsiteConf("vimm.json")
	// specificInfo := config.SpecificInfo
	itemKeys := config.Search.ItemKeys

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {

		item := variables.Item{
			Name: h.ChildText(itemKeys.Name),
			Link: h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
		}
		// item.Metadata["console"] = h.ChildText(specificInfo["console"])

		// returns the right link but the browser can't load it properly -> vimm responds with a default image instead
		item.Thumbnail = "https://vimm.net/image.php?type=box&id=" + strings.Split(item.Link, "/vault/")[1]

		results = append(results, item)
	})

	c.Visit(config.Search.Url + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement))

	return results
}

// func getThumbnail(link string, cssSelector string, attribute string) string {
// 	c := colly.NewCollector()
// 	var thumbnailLink string
// 	c.OnHTML(cssSelector, func(h *colly.HTMLElement) {
// 		thumbnailLink = "https://vimm.net" + h.Attr(attribute)
// 		fmt.Println(h.Attr(attribute))
// 	})

// 	c.Visit(link)

// 	return thumbnailLink
// }
