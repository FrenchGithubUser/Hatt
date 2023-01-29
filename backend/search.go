package main

import (
	"encoding/json"
	"fmt"
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

func deserializeWebsiteConf(file string) Config {
	var config Config

	content, err := ioutil.ReadFile("./website_configs/" + file)
	if err != nil {
		fmt.Println("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &config)

	return config
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
	fmt.Println(config.Name)

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

	configFiles, _ := ioutil.ReadDir("./website_configs")
	for _, configFile := range configFiles {
		var conf Config = deserializeWebsiteConf(configFile.Name())
		for _, category := range categories {
			if websiteHasCategory(conf.Categories, category) {
				configs = append(configs, conf)
				break
			}
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
