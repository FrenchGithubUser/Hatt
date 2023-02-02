package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func deserializeWebsiteConf(file string, dev bool) Config {
	var config Config

	var configs_dir string
	if dev {
		configs_dir = CONFIGS_DIR + "dev/"
	} else {
		configs_dir = CONFIGS_DIR
	}

	content, err := ioutil.ReadFile(configs_dir + file)
	if err != nil {
		fmt.Println("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &config)

	return config
}

func deserializeCredentials() Credentials {
	var credentials Credentials = map[string]map[string]map[string]string{}

	credsList, _ := ioutil.ReadFile(credentialsPath)
	json.Unmarshal(credsList, &credentials)

	return credentials
}
