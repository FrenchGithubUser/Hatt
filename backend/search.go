package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gocolly/colly"
)

type item struct {
	Name      string
	Thumbnail string
	Link      string
}

func deserializeWebsiteConf(website string) Config {
	var config Config

	content, err := ioutil.ReadFile("./website_configs/" + website + ".json")
	if err != nil {
		fmt.Println("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &config)

	return config
}

func getWebsiteData(input string, website string) []item {

	var config Config = deserializeWebsiteConf(website)

	var items []item
	c := colly.NewCollector()
	itemKeys := config.Search.ItemKeys

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {

		// in case an html tag doesn't have any content
		if h.ChildText(itemKeys.Name) != "" {

			item := item{
				Name:      h.ChildText(itemKeys.Name),
				Thumbnail: h.ChildAttr(itemKeys.Thumbnail.Key, itemKeys.Thumbnail.Attribute),
				Link:      h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
			}

			items = append(items, item)
		}
	})

	// pagination handling
	// c.OnHTML("a.navigation.next", func(h *colly.HTMLElement) {
	// 	nextPage := h.Request.AbsoluteURL(h.Attr("href"))
	// 	c.Visit(nextPage)

	// })

	c.Visit(config.Search.Url + input)

	return items
}

func getItemsList(w http.ResponseWriter, r *http.Request) {

	input := r.URL.Query().Get("input")

	var websites []string = []string{"f2movies"}

	var results []item

	for _, website := range websites {
		items := getWebsiteData(input, website)
		results = append(results, items...)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(results)
}
