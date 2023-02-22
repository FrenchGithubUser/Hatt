package main

import (
	"hatt/configuration"
	"hatt/helpers"
	"hatt/variables"
	"io/ioutil"
)

func (a *App) GetWebsites(selectedCategories map[string][]string) []string {
	categories := selectedCategories["categories"]
	var websites []string
	configFiles, _ := ioutil.ReadDir(variables.CONFIGS_DIR)
	for _, configFile := range configFiles {
		var conf configuration.Config = helpers.DeserializeWebsiteConf(configFile.Name())
		for _, category := range categories {
			if helpers.WebsiteHasCategory(conf.Categories, category) {
				websites = append(websites, conf.Name)
				break
			}
		}
	}
	return websites
}
