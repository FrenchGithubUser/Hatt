package specificScrapers

import (
	"hatt/assets"
	"hatt/helpers"
	"hatt/login"
	"hatt/variables"
	"strings"

	"github.com/gocolly/colly"
)

func (t T) Mobilism() []variables.Item {

	var results []variables.Item
	c := colly.NewCollector()

	config := assets.DeserializeWebsiteConf("mobilism.json")
	itemKeys := config.Search.ItemKeys

	login.Login("mobilism")

	c.OnHTML("#message", func(h *colly.HTMLElement) {
		if strings.Contains(h.ChildText("p"), "Sorry but you are not permitted to use the search system") {
			message := variables.Item{
				Name: "error",
				Metadata: map[string]string{
					"name": "login_required",
				},
			}
			results = append(results, message)
		}
	})

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {

		item := variables.Item{
			Name:      h.ChildText(itemKeys.Name),
			Thumbnail: "",
			Link:      h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
		}

		if item.Name != "" {

			item.Metadata = map[string]string{
				"category": h.ChildText("a:last-of-type"),
			}

			results = append(results, item)
		}
	})

	tokens := helpers.DeserializeCredentials("mobilism").Tokens
	c.OnRequest(func(r *colly.Request) {
		var tokensString string
		for tokenName, token := range tokens {
			tokensString += tokenName + "=" + token["value"]
		}
		r.Headers.Set("Cookie", tokensString)
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:109.0) Gecko/20100101 Firefox/109.0")
	})

	c.Visit(config.Search.Url + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement))

	return results
}
