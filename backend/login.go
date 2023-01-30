package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func login(website string) string {
	var conf Config = deserializeWebsiteConf(website + ".json")

	savedCredentials := map[string]map[string]string{}

	credsList, _ := ioutil.ReadFile("../settings/credentials.json")
	json.Unmarshal(credsList, &savedCredentials)
	websiteCredentials := savedCredentials[website]

	hc := http.Client{}

	form := url.Values{}
	for field, value := range websiteCredentials {
		form.Add(field, value)
	}

	req, _ := http.NewRequest("POST", conf.Login.Url, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)

	return ""
}
