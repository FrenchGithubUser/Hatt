package main

import (
	"encoding/json"
	"fmt"
	"hatt/configuration"
	"hatt/specificScrapers"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
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

func getItemsList(w http.ResponseWriter, r *http.Request) {

	input := r.URL.Query().Get("input")

	configs := []configuration.Config{}

	if ENV == "dev" {
		configFiles, _ := ioutil.ReadDir(CONFIGS_DIR + "/dev")
		for _, configFile := range configFiles {
			var conf configuration.Config = deserializeWebsiteConf(configFile.Name())
			configs = append(configs, conf)
			fmt.Println(conf.Name)
		}
	} else {
		categories := strings.Split(r.URL.Query().Get("categories"), ",")
		configFiles, _ := ioutil.ReadDir(CONFIGS_DIR)
		for _, configFile := range configFiles {
			var conf configuration.Config = deserializeWebsiteConf(configFile.Name())
			for _, category := range categories {
				if websiteHasCategory(conf.Categories, category) {
					configs = append(configs, conf)
					break
				}
			}
		}
	}

	var results []itemList
	for _, config := range configs {
		if config.SpecificScraper {
			t := specificScrapers.T{}
			specificFunction := reflect.ValueOf(t).MethodByName(strings.Title(config.Name))
			specificFunction.Call(nil)
		}
		items := scrapePlainHtml(input, config)
		result := itemList{
			Website: config.Name,
			Items:   items,
		}
		results = append(results, result)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(results)
}
