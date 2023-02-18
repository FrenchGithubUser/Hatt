package main

import (
	"encoding/json"
	"fmt"
	"hatt/configuration"
	"hatt/helpers"
	"hatt/htmlParsers"
	"hatt/specificScrapers"
	specificScrapersDev "hatt/specificScrapers/dev"
	"hatt/variables"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"sync"
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

	variables.RESULTS = []variables.ItemList{}
	var wg sync.WaitGroup
	for _, config := range configs {
		var items []variables.Item
		var specificFunction reflect.Value
		wg.Add(1)
		go func(config configuration.Config) {
			fmt.Println("starting with : " + config.Name)
			if config.SpecificScraper {
				if variables.ENV == "dev" {
					t := specificScrapersDev.T{}
					specificFunction = reflect.ValueOf(t).MethodByName(strings.Title(config.Name))
				} else {
					t := specificScrapers.T{}
					specificFunction = reflect.ValueOf(t).MethodByName(strings.Title(config.Name))
				}

				items = specificFunction.Call(nil)[0].Interface().([]variables.Item)

			} else {
				items = htmlParsers.ScrapePlainHtml(config)
			}
			result := variables.ItemList{
				Website: config.Name,
				Items:   items,
			}
			variables.RESULTS = append(variables.RESULTS, result)
			wg.Done()
			fmt.Println("done with : " + config.Name)
		}(config)
	}

	wg.Wait()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(variables.RESULTS)
}
