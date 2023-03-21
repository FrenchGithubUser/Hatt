package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"hatt/variables"
	"io/ioutil"
)

// type Credentials = map[string]map[string]map[string]map[string]string

type Helper struct {
	ctx context.Context
}

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

func GetAllWebsiteCredentials() Credentials {
	var credentials Credentials

	credsList, err := ioutil.ReadFile(variables.CREDENTIALS_PATH)
	if err != nil {
		//todo : make sure that the error is "no such file or directory (os.patherror ?)"
		creationErr := ioutil.WriteFile(variables.CREDENTIALS_PATH, []byte("{}"), 0755)
		if creationErr != nil {
			fmt.Println(creationErr)
		}
		credsList, _ = ioutil.ReadFile(variables.CREDENTIALS_PATH)
	}
	json.Unmarshal(credsList, &credentials)

	return credentials
}

func (h *Helper) DeserializeCredentials(website string) WebsiteCredentials {

	credentials := GetAllWebsiteCredentials()

	var websiteCredentials WebsiteCredentials
	for _, siteCreds := range credentials {
		if siteCreds.Name == website {
			websiteCredentials = siteCreds
		}
	}

	return websiteCredentials
}

func (h *Helper) SaveUpdatedCredentials(site string, updatedCredentials WebsiteCredentials) {
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

func FormatDuration(duration int) string {
	hours := duration / 3600 // integer division, decimals are truncated
	minutes := (duration - hours*3600) / 60
	seconds := duration - hours*3600 - minutes*60

	minutesExtraZero := ""
	secondsExtraZero := ""
	if minutes < 10 {
		minutesExtraZero = "0"
	}
	if seconds < 10 {
		secondsExtraZero = "0"
	}

	var formattedDuration string
	if hours != 0 {
		formattedDuration = fmt.Sprintf("%d:", hours)
	}
	formattedDuration += fmt.Sprintf("%s%d:%s%d", minutesExtraZero, minutes, secondsExtraZero, seconds)

	return formattedDuration

}
