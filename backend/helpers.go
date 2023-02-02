package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func deserializeWebsiteConf(file string) Config {
	var config Config

	content, err := ioutil.ReadFile("./website_configs/" + file)
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

	fmt.Println(credentials)

	return credentials
}
