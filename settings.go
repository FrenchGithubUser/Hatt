package main

import (
	"encoding/json"
	"fmt"
	"hatt/assets"
	"hatt/variables"
	"io/ioutil"
)

type UserSettings struct {
	Appearance Appearance `json:"appearance"`
}

type Appearance struct {
	ThumbnailsSize int `json:"thumbnailsSize"`
}

func (a *App) ReadUserSettings() UserSettings {
	var settings UserSettings
	var settingsFile []byte
	settingsFile, err := ioutil.ReadFile(variables.SETTINGS_PATH)
	fmt.Println(err)
	if err != nil {
		assets.CopyBaseSettings()
		settingsFile, _ = ioutil.ReadFile(variables.SETTINGS_PATH)
	}
	json.Unmarshal(settingsFile, &settings)

	return settings
}
