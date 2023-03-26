package htmlParsers

import (
	"fmt"
	"hatt/configuration"
	"hatt/helpers"
	"hatt/variables"
	"strings"

	"github.com/gocolly/colly"
)

func ScrapePlainHtml(config configuration.Config) []variables.Item {

	var items []variables.Item
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; rv:109.0) Gecko/20100101 Firefox/109.0"
	itemKeys := config.Search.ItemKeys

	// c.OnHTML("body", func(h *colly.HTMLElement) {
	// 	fmt.Println(h)
	// })

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {
		item := variables.Item{
			Name: h.ChildText(itemKeys.Name),
		}

		if itemKeys.Link == "root" {
			item.Link = h.Request.AbsoluteURL(h.Attr("href"))
		} else {
			item.Link = h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href"))
		}

		var thumbnailLink string
		if itemKeys.Thumbnail.Key == "root" {
			thumbnailLink = h.Attr(itemKeys.Thumbnail.Attribute)
		} else {
			thumbnailLink = h.ChildAttr(itemKeys.Thumbnail.Key, itemKeys.Thumbnail.Attribute)
		}
		if itemKeys.Thumbnail.AppendToSiteUrl {
			item.Thumbnail = h.Request.AbsoluteURL(thumbnailLink)
		} else {
			item.Thumbnail = thumbnailLink
		}

		item.Metadata = map[string]string{}
		for metadata, key := range itemKeys.Metadata {
			info := h.ChildText(key)
			if info != "" {
				item.Metadata[metadata] = info
			}
		}

		if item.Name != "" {
			items = append(items, item)
		}
	})

	// pagination handling
	// c.OnHTML("a.navigation.next", func(h *colly.HTMLElement) {
	// 	nextPage := h.Request.AbsoluteURL(h.Attr("href"))
	// 	c.Visit(nextPage)

	// })

	// when website requires login
	// if config.Login.Url != "" {
	// 	// login(config.Name)
	// 	// tokens := helpers.DeserializeCredentials()[config.Name]["tokens"]
	// 	c.OnRequest(func(r *colly.Request) {
	// 		// for _, token := range config.Login.Tokens {
	// 		// }
	// 		r.Headers.Set("cookie", "")
	// 	})
	// }

	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println(r)
	// })

	// c.OnError(func(r *colly.Response, err error) { fmt.Println(r, err) })

	c.OnRequest(func(r *colly.Request) {
		// maybe set this value according to the user's locale ?
		r.Headers.Set("Accept-Language", "en")
	})

	if config.Search.Method == "POST" {
		formData := map[string]string{}
		for key, value := range config.Search.PostFields.Generic {
			formData[key] = value
		}
		formData[config.Search.PostFields.Input] = variables.CURRENT_INPUT
		c.Post(config.Search.Url, formData)
	} else {
		formattedInput := config.Search.Url + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement)
		// some websites support multiple categories, to avoid irrelevant results, only search in relevant categories
		if config.Search.CategorySpecificAttributes.Name != "" {
			for category, value := range config.Search.CategorySpecificAttributes.Values {
				// searches for all the categories if no category is selected (that happens if the source is in a custom list)
				if helpers.IsStringInSlice(variables.SELECTED_CATEGORIES, category) || len(variables.SELECTED_CATEGORIES) == 0 {
					fmt.Println(formattedInput + "&" + config.Search.CategorySpecificAttributes.Name + "=" + value)
					c.Visit(formattedInput + "&" + config.Search.CategorySpecificAttributes.Name + "=" + value)
				}
			}
		} else {
			c.Visit(formattedInput)
		}
	}

	return items
}
