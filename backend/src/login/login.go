package login

import (
	"fmt"
	"hatt/configuration"
	"hatt/helpers"
	"net/http"
	"net/url"
	"strings"
)

func Login(website string) {
	var conf configuration.Config = helpers.DeserializeWebsiteConf(website + ".json")

	var websiteCredentials helpers.WebsiteCredentials
	websiteCredentials = helpers.DeserializeCredentials(website)

	hc := http.Client{}

	form := url.Values{}
	for field, value := range websiteCredentials.LoginInfo {
		form.Add(field, value)
	}

	req, _ := http.NewRequest("POST", conf.Login.Url, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:109.0) Gecko/20100101 Firefox/109.0")

	resp, err := hc.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	websiteCredentials.Tokens = make([]helpers.Token, 0)
	for _, cookie := range resp.Cookies() {
		token := helpers.Token{
			Name:    cookie.Name,
			Value:   cookie.Value,
			Expires: cookie.Expires.String(),
		}
		websiteCredentials.Tokens = append(websiteCredentials.Tokens, token)
	}

	helpers.SaveUpdatedCredentials(website, websiteCredentials)

}
