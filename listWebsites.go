package main

import (
	"hatt/assets"
	"hatt/configuration"
	"hatt/helpers"
)

func (a *App) GetWebsitesWithCategories(selectedCategories map[string][]string) []string {

	categories := selectedCategories["categories"]
	var websites []string
	configFiles := assets.GetWebsiteConfigs()
	for _, configFile := range configFiles {
		var conf configuration.Config = assets.DeserializeWebsiteConf(configFile.Name())
		for _, category := range categories {
			if helpers.WebsiteHasCategory(conf.Categories, category) || category == "all" {
				websites = append(websites, conf.Name)
				break
			}
		}
	}
	return websites
}

type website struct {
	Name   string
	Fields []string
}

func (a *App) GetWebsitesWithLogin() []website {
	var websites []website
	configFiles := assets.GetWebsiteConfigs()
	for _, configFile := range configFiles {
		var config configuration.Config = assets.DeserializeWebsiteConf(configFile.Name())
		if config.Login.Url != "" {
			websites = append(websites, website{
				Name:   config.Name,
				Fields: config.Login.Fields,
			})
		}
	}
	return websites
}
