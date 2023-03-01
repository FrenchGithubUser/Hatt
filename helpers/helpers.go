package helpers

import (
	"encoding/json"
	"fmt"
	"hatt/variables"
	"io/ioutil"
)

// type Credentials = map[string]map[string]map[string]map[string]string

type Credentials = []WebsiteCredentials

type WebsiteCredentials struct {
	Name      string
	LoginInfo map[string]string
	Tokens    map[string]map[string]string
}

func WebsiteHasCategory(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// func DeserializeWebsiteConf(file string) configuration.Config {

// 	var config configuration.Config

// 	var configs_dir string
// 	if variables.ENV == "dev" {
// 		configs_dir = variables.CONFIGS_DIR + "dev/"
// 	} else {
// 		configs_dir = variables.CONFIGS_DIR
// 	}

// 	content, err := ioutil.ReadFile(configs_dir + file)
// 	if err != nil {
// 		fmt.Println("Error when opening file: ", err)
// 	}

// 	err = json.Unmarshal(content, &config)

// 	return config
// }

func DeserializeCredentials(website string) WebsiteCredentials {
	var credentials Credentials

	credsList, err := ioutil.ReadFile(variables.CREDENTIALS_PATH)
	if err != nil {
		//todo make sure that the error is "no such file or directory (os.patherror ?)"
		creationErr := ioutil.WriteFile(variables.CREDENTIALS_PATH, []byte("{}"), 0755)
		if creationErr != nil {
			fmt.Println(creationErr)
		}
		credsList, _ = ioutil.ReadFile(variables.CREDENTIALS_PATH)
	}
	json.Unmarshal(credsList, &credentials)

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
	oldCredentials, _ := ioutil.ReadFile(variables.CREDENTIALS_PATH)
	json.Unmarshal(oldCredentials, &credentials)

	var i int = 0
	for _, websiteCredentials := range credentials {
		if websiteCredentials.Name == site {
			credentials[i] = updatedCredentials
		}
		i++
	}

	updatedCredentialsJson, _ := json.Marshal(credentials)
	_ = ioutil.WriteFile(variables.CREDENTIALS_PATH, updatedCredentialsJson, 0644)
}
