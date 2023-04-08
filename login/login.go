package login

import (
	"fmt"
	"hatt/assets"
	"hatt/configuration"
	"hatt/helpers"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func Login(website string) {
	var conf configuration.Config = assets.DeserializeWebsiteConf(website + ".json")

	var websiteCredentials helpers.WebsiteCredentials

	h := &helpers.Helper{}
	websiteCredentials = h.DeserializeCredentials(website)

	// if struct is empty, credentials were not given by the user
	if websiteCredentials.Name == "" {
		return
	}

	for _, confCookieName := range conf.Login.Tokens {
		cookieExpirationDate, _ := strconv.Atoi(websiteCredentials.Tokens[confCookieName]["expires"])
		now := int(time.Now().UnixMilli())

		// checking if cookie is expired or will expire in the next 20 seconds (small error margin)
		isCookieExpired := cookieExpirationDate-now < 20

		fmt.Println(now, cookieExpirationDate, isCookieExpired)

		if !isCookieExpired {
			return
		}
	}

	websiteCredentials.Tokens = map[string]map[string]string{}

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

	for _, cookie := range resp.Cookies() {
		for _, confCookieName := range conf.Login.Tokens {
			if confCookieName == cookie.Name {
				token := map[string]string{
					"value":   cookie.Value,
					"expires": fmt.Sprint(time.Now().Add(time.Duration(conf.Login.TokenLifespan) * time.Second).UnixMilli()), //cookie.Expires.String(),
				}
				websiteCredentials.Tokens[cookie.Name] = token
			}
		}
	}

	h.SaveUpdatedCredentials(website, websiteCredentials)

}
