package main

import (
	"encoding/json"
	"hatt/assets"
	"hatt/variables"
	"io/ioutil"
)

type UserSettings struct {
	General General `json:"general"`
}

type General struct {
	ThumbnailsSize int    `json:"thumbnailsSize"`
	Lang           string `json:"lang"`
}

func (a *App) ReadUserSettings() UserSettings {
	var settings UserSettings
	var settingsFile []byte
	settingsFile, err := ioutil.ReadFile(variables.SETTINGS_PATH)
	if err != nil {
		assets.CopyBaseSettings()
		settingsFile, _ = ioutil.ReadFile(variables.SETTINGS_PATH)
	}
	json.Unmarshal(settingsFile, &settings)

	return settings
}

func (a *App) UpdateUserSettings(updatedSettings UserSettings) {
	updatedSettingsFile, _ := json.Marshal(updatedSettings)
	_ = ioutil.WriteFile(variables.SETTINGS_PATH, updatedSettingsFile, 0644)

}
