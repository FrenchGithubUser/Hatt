package helpers

import (
	"encoding/json"
	"fmt"
	"hatt/configuration"
	"hatt/variables"
	"io/ioutil"
)

type Credentials = map[string]map[string]map[string]string

func DeserializeWebsiteConf(file string) configuration.Config {
	var config configuration.Config

	var configs_dir string
	if variables.ENV == "dev" {
		configs_dir = variables.CONFIGS_DIR + "dev/"
	} else {
		configs_dir = variables.CONFIGS_DIR
	}

	content, err := ioutil.ReadFile(configs_dir + file)
	if err != nil {
		fmt.Println("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &config)

	return config
}

func DeserializeCredentials() Credentials {
	var credentials Credentials = map[string]map[string]map[string]string{}

	credsList, _ := ioutil.ReadFile(variables.CREDENTIALS_PATH)
	json.Unmarshal(credsList, &credentials)

	return credentials
}
