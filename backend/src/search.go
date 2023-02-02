package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

type item struct {
	Name      string
	Thumbnail string
	Link      string
}

type itemList struct {
	Website string
	Items   []item
}

func websiteHasCategory(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func getWebsiteData(input string, config Config) []item {

	var items []item
	c := colly.NewCollector()
	itemKeys := config.Search.ItemKeys

	c.OnHTML(itemKeys.Root, func(h *colly.HTMLElement) {

		item := item{
			Name:      h.ChildText(itemKeys.Name),
			Thumbnail: h.ChildAttr(itemKeys.Thumbnail.Key, itemKeys.Thumbnail.Attribute),
			Link:      h.Request.AbsoluteURL(h.ChildAttr(itemKeys.Link, "href")),
		}

		items = append(items, item)
	})

	// pagination handling
	// c.OnHTML("a.navigation.next", func(h *colly.HTMLElement) {
	// 	nextPage := h.Request.AbsoluteURL(h.Attr("href"))
	// 	c.Visit(nextPage)

	// })

	c.Visit(config.Search.Url + strings.ReplaceAll(input, " ", config.Search.SpaceReplacement))

	return items
}

func getItemsList(w http.ResponseWriter, r *http.Request) {

	input := r.URL.Query().Get("input")
	categories := strings.Split(r.URL.Query().Get("categories"), ",")

	configs := []Config{}

	configFiles, _ := ioutil.ReadDir(CONFIGS_DIR)
	for _, configFile := range configFiles {
		var conf Config = deserializeWebsiteConf(configFile.Name(), false)
		for _, category := range categories {
			if websiteHasCategory(conf.Categories, category) {
				configs = append(configs, conf)
				break
			}
		}
	}
	if ENV == "dev" {
		configFiles, _ := ioutil.ReadDir(CONFIGS_DIR + "/dev")
		for _, configFile := range configFiles {
			var conf Config = deserializeWebsiteConf(configFile.Name(), true)
			configs = append(configs, conf)
		}
	}

	var results []itemList
	for _, config := range configs {
		items := getWebsiteData(input, config)
		result := itemList{
			Website: config.Name,
			Items:   items,
		}
		results = append(results, result)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(results)
}
