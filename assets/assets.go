package assets

import (
	"embed"
	"encoding/json"
	"fmt"
	"hatt/configuration"
	"io/fs"
	"os"
)

type Credentials = []WebsiteCredentials

type WebsiteCredentials struct {
	Name      string
	LoginInfo map[string]string
	Tokens    map[string]map[string]string
}

//go:embed website_configs
var WebsiteConfigs embed.FS

//go:embed settings/credentials.json
var WebsiteCreds embed.FS

func GetWebsiteConfigs() []fs.DirEntry {

	directory, err := WebsiteConfigs.ReadDir("website_configs")

	if err != nil {
		panic(err)
	}

	return directory

}

func DeserializeWebsiteConf(filename string) configuration.Config {

	var config configuration.Config
	content, _ := WebsiteConfigs.ReadFile("website_configs/" + filename)
	json.Unmarshal([]byte(content), &config)
	return config
}

func DeserializeCredentials(website string) WebsiteCredentials {
	var credentials Credentials

	content, err := WebsiteCreds.ReadFile("settings/credentials.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(content, &credentials)

	var websiteCredentials WebsiteCredentials
	for _, siteCreds := range credentials {
		if siteCreds.Name == website {
			websiteCredentials = siteCreds
		}
	}

	return websiteCredentials
}

func SaveUpdatedCredentials(site string, updatedCredentials WebsiteCredentials) {
	var credentials Credentials
	oldCredentials, _ := WebsiteCreds.ReadFile("settings/credentials.json")
	json.Unmarshal(oldCredentials, &credentials)

	var i int = 0
	for _, websiteCredentials := range credentials {
		if websiteCredentials.Name == site {
			credentials[i] = updatedCredentials
		}
		i++
	}

	updatedCredentialsJson, _ := json.Marshal(credentials)
	_ = os.WriteFile("settings/credentials.json", updatedCredentialsJson, 0644)
}
