package main

import (
	"encoding/json"
	"fmt"
	"hatt/assets"
	"hatt/variables"
	"io/ioutil"
)

type UserSettings struct {
	General General `json:"general"`
}

type General struct {
	ThumbnailsSize    int    `json:"thumbnailsSize"`
	Lang              string `json:"lang"`
	DarkMode          bool   `json:"darkMode"`
	ItemClickedAction string `json:"itemClickedAction"`
}

type CustomList struct {
	Name    string   `json:"name"`
	Sources []string `json:"sources"`
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

func (a *App) ReadCustomLists() []CustomList {
	var customLists []CustomList
	var customListsFile []byte
	customListsFile, err := ioutil.ReadFile(variables.CUSTOM_LISTS_PATH)
	if err != nil {
		creationErr := ioutil.WriteFile(variables.CUSTOM_LISTS_PATH, []byte("[]"), 0755)
		if creationErr != nil {
			fmt.Println("Error when creating custom lists file : ", creationErr)
		}
		customListsFile, _ = ioutil.ReadFile(variables.SETTINGS_PATH)
	}
	json.Unmarshal(customListsFile, &customLists)

	return customLists
}

func (a *App) UpdateCustomLists(updatedLists []CustomList) {
	updatedListsFile, _ := json.Marshal(updatedLists)
	_ = ioutil.WriteFile(variables.CUSTOM_LISTS_PATH, updatedListsFile, 0644)
}
