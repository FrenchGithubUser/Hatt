package main

import (
	"encoding/json"
	"fmt"
	"hatt/configuration"
	"hatt/helpers"
	"hatt/specificScrapers"
	specificScrapersDev "hatt/specificScrapers/dev"
	"hatt/variables"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func websiteHasCategory(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func getItemsList(w http.ResponseWriter, r *http.Request) {

	variables.CURRENT_INPUT = r.URL.Query().Get("input")

	configs := []configuration.Config{}

	if variables.ENV == "dev" {
		configFiles, _ := ioutil.ReadDir(variables.CONFIGS_DIR + "/dev")
		for _, configFile := range configFiles {
			var conf configuration.Config = helpers.DeserializeWebsiteConf(configFile.Name())
			fmt.Println(conf.Name)
			configs = append(configs, conf)
		}
	} else {
		categories := strings.Split(r.URL.Query().Get("categories"), ",")
		configFiles, _ := ioutil.ReadDir(variables.CONFIGS_DIR)
		for _, configFile := range configFiles {
			var conf configuration.Config = helpers.DeserializeWebsiteConf(configFile.Name())
			for _, category := range categories {
				if websiteHasCategory(conf.Categories, category) {
					configs = append(configs, conf)
					break
				}
			}
		}
	}

	var results []variables.ItemList
	for _, config := range configs {
		var items []variables.Item
		if config.SpecificScraper {
			if variables.ENV == "dev" {
				t := specificScrapersDev.T{}
				specificFunction := reflect.ValueOf(t).MethodByName(strings.Title(config.Name))
				items = specificFunction.Call(nil)[0].Interface().([]variables.Item)
			} else {
				t := specificScrapers.T{}
				specificFunction := reflect.ValueOf(t).MethodByName(strings.Title(config.Name))
				items = specificFunction.Call(nil)[0].Interface().([]variables.Item)
			}
		} else {
			items = scrapePlainHtml(config)
		}
		result := variables.ItemList{
			Website: config.Name,
			Items:   items,
		}
		results = append(results, result)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(results)
}
