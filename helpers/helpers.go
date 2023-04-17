package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"hatt/configuration"
	"hatt/variables"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/go-rod/rod/lib/launcher"
	"github.com/gocolly/colly"
)

type Helper struct {
	ctx context.Context
}

type Credentials = []WebsiteCredentials

type WebsiteCredentials struct {
	Name      string
	LoginInfo map[string]string
	Tokens    map[string]map[string]string
}

func IsStringInSlice(s []string, str string) bool {
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
	e := json.Unmarshal(credsList, &credentials)

	if e != nil {
		fmt.Println("Error when deserializing credentials", e)
	}

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

	// desserialize the file which has credentials from all the websites
	oldCredentials, _ := ioutil.ReadFile(variables.CREDENTIALS_PATH)
	json.Unmarshal(oldCredentials, &credentials)

	updatedCredentials.Name = site

	// remove old creds if they exist
	var credentialsIndex int = -1
	for i, websiteCredentials := range credentials {
		if websiteCredentials.Name == site {
			credentialsIndex = i
			break
		}
	}
	if credentialsIndex > -1 {
		credentials[credentialsIndex] = credentials[len(credentials)-1]
		credentials = credentials[:len(credentials)-1]
	}

	// add new credentials
	// for key, value := range assets.DeserializeWebsiteConf(site + ".json").Login.GenericFields {
	// 	updatedCredentials.LoginInfo[key] = value
	// }

	// add updated credentials of the site in the list of WebsiteCredentials
	credentials = append(credentials, updatedCredentials)

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

// returns true if the config is not Xxx, or if the user enabled Xxx
func XxxCheck(config configuration.Config) bool {
	fmt.Println(config.Xxx)
	return (!config.Xxx || variables.CURRENT_USER_SETTINGS.General.Xxx)
}

func GetServerGeneratedTokens(url string, tokenNames []string) map[string]string {
	c := colly.NewCollector()
	// add all headers, not just user agent

	tokens := map[string]string{}

	c.OnHTML("input", func(h *colly.HTMLElement) {
		for _, tokenName := range tokenNames {
			if h.Attr("name") == tokenName {
				tokens[tokenName] = h.Attr("value")
			}
		}
	})

	c.Visit(url)

	return tokens
}

func IsLoginNeeded(websiteCredentials WebsiteCredentials, conf configuration.Config) bool {
	var loginNeeded bool = false
	for _, confCookieName := range conf.Login.Tokens {
		cookieExpirationDate, _ := strconv.Atoi(websiteCredentials.Tokens[confCookieName]["expires"])
		now := int(time.Now().UnixMilli())

		// checking if cookie is expired or will expire in the next 20 seconds (small error margin)
		isCookieExpired := cookieExpirationDate-now < 20

		if isCookieExpired {
			loginNeeded = true
		}
	}
	return loginNeeded
}

func InstanciateBrowser() string {
	// todo : change executable paths based on the OS
	var l string
	var browserErr error
	l, browserErr = launcher.New().Bin("/usr/bin/brave").Launch()
	if browserErr != nil {
		l, browserErr = launcher.New().Bin("/usr/bin/chromium").Launch()
	}
	if browserErr != nil {
		l, browserErr = launcher.New().Bin("/usr/bin/chrome").Launch()
	}

	return l
}
