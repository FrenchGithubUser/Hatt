package assets

import (
	"embed"
	"encoding/json"
	"fmt"
	"hatt/configuration"
	"hatt/variables"
	"io/fs"
	"io/ioutil"
)

//go:embed website_configs
var WebsiteConfigs embed.FS

// //go:embed settings/credentials.json
// var WebsiteCreds embed.FS

//go:embed compatible_downloaders.json
var compatibleDownlodersFile embed.FS

//go:embed base_settings.json
var baseSettings embed.FS

func GetWebsiteConfigs() []fs.DirEntry {

	directory, err := WebsiteConfigs.ReadDir(variables.CONFIGS_DIR)

	if err != nil {
		panic(err)
	}

	return directory

}

func DeserializeWebsiteConf(filename string) configuration.Config {

	var config configuration.Config

	content, err := WebsiteConfigs.ReadFile(variables.CONFIGS_DIR + "/" + filename)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(content), &config)

	return config
}

func InitCompatibleDownloaders() {

	content, err := compatibleDownlodersFile.ReadFile("compatible_downloaders.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(content), &variables.CompatibleDownloaders)

}

func CopyBaseSettings() {
	content, err := baseSettings.ReadFile("base_settings.json")
	if err != nil {
		fmt.Println(err)
	}
	_ = ioutil.WriteFile(variables.SETTINGS_PATH, content, 0644)
	// json.Unmarshal([]byte(content), &variables.CompatibleDownloaders)
}

// func DeserializeCredentials(website string) WebsiteCredentials {
// 	var credentials Credentials

// 	content, err := WebsiteCreds.ReadFile("settings/credentials.json")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	json.Unmarshal(content, &credentials)

// 	var websiteCredentials WebsiteCredentials
// 	for _, siteCreds := range credentials {
// 		if siteCreds.Name == website {
// 			websiteCredentials = siteCreds
// 		}
// 	}

// 	return websiteCredentials
// }

// func SaveUpdatedCredentials(site string, updatedCredentials WebsiteCredentials) {
// 	var credentials Credentials
// 	oldCredentials, _ := WebsiteCreds.ReadFile("settings/credentials.json")
// 	json.Unmarshal(oldCredentials, &credentials)

// 	var i int = 0
// 	for _, websiteCredentials := range credentials {
// 		if websiteCredentials.Name == site {
// 			credentials[i] = updatedCredentials
// 		}
// 		i++
// 	}

// 	updatedCredentialsJson, _ := json.Marshal(credentials)
// 	_ = os.WriteFile("settings/credentials.json", updatedCredentialsJson, 0644)
// }
