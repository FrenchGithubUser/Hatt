package main

import (
	"hatt/assets"
	"hatt/configuration"
	"hatt/helpers"
)

func (a *App) GetWebsites(selectedCategories map[string][]string) []string {

	categories := selectedCategories["categories"]
	var websites []string
	configFiles := assets.GetWebsiteConfigs() //ioutil.ReadDir(variables.CONFIGS_DIR)
	for _, configFile := range configFiles {
		var conf configuration.Config = assets.DeserializeWebsiteConf(configFile.Name())
		for _, category := range categories {
			if helpers.WebsiteHasCategory(conf.Categories, category) {
				websites = append(websites, conf.Name)
				break
			}
		}
	}
	return websites
}
