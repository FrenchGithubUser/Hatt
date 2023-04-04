package specificScrapers

import (
	"hatt/assets"
	"hatt/variables"
	"strings"

	"github.com/gocolly/colly"
)

func (t T) Vimm() []variables.Item {

	var results []variables.Item
	c := colly.NewCollector()

	config := assets.DeserializeWebsiteConf("vimm.json")
	specificInfo := config.SpecificInfo
	itemKeys := config.Search.ItemKeys

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {
		item := variables.Item{
			Name: h.ChildText(itemKeys.Name),
			Link: h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
		}
		// a cookie is needed to load the image, otherwise vimm returns a default image
		// item.Thumbnail = "https://vimm.net/image.php?type=box&id=" + strings.Split(item.Link, "/vault/")[1]

		if item.Name != "" {
			item.Metadata = map[string]string{
				"console": h.ChildText(specificInfo["console"]),
			}
			//sometimes the region is not referenced
			flag := strings.Split(h.ChildAttr(specificInfo["region"], "src"), "/flags/")
			if len(flag) >= 2 {
				item.Metadata["region"] = strings.ReplaceAll(flag[1], ".png", "")
			}
			results = append(results, item)
		}
	})

	c.Visit(config.Search.Url + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement))

	return results
}
