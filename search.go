package main

import (
	"fmt"
	"hatt/configuration"
	"hatt/helpers"
	"hatt/htmlParsers"
	"hatt/specificScrapers"
	specificScrapersDev "hatt/specificScrapers/dev"
	"hatt/variables"
	"io/ioutil"
	"reflect"
	"strings"
	"sync"
)

func (a *App) Search(userInput string, websites []string) []variables.ItemList {
	fmt.Println(websites)

	variables.CURRENT_INPUT = userInput

	configs := []configuration.Config{}

	if variables.ENV == "dev" {
		configFiles, _ := ioutil.ReadDir(variables.CONFIGS_DIR + "dev")
		for _, configFile := range configFiles {
			var conf configuration.Config = helpers.DeserializeWebsiteConf(configFile.Name())
			configs = append(configs, conf)
		}
	} else {
		for _, website := range websites {
			var conf configuration.Config = helpers.DeserializeWebsiteConf(website + ".json")
			configs = append(configs, conf)
		}
	}
	fmt.Println(configs)

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

	return variables.RESULTS

}
