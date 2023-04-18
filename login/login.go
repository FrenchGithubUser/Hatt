package login

import (
	"fmt"
	"hatt/assets"
	"hatt/configuration"
	"hatt/helpers"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Login(website string) bool /*return whether or not the login was successfull*/ {
	var conf configuration.Config = assets.DeserializeWebsiteConf(website + ".json")

	var websiteCredentials helpers.WebsiteCredentials

	h := &helpers.Helper{}
	websiteCredentials = h.DeserializeCredentials(website)

	if websiteCredentials.Name == "" {
		// no credentials provided
		return false
	}

	isLoginNeeded := helpers.IsLoginNeeded(websiteCredentials, conf)

	if !isLoginNeeded {
		fmt.Println(isLoginNeeded)
		return true
	}

	websiteCredentials.Tokens = map[string]map[string]string{}

	hc := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// todo : experiment with this to get the cookies, even after a redirect
	// hc.Jar.Cookies(&url.URL{})

	form := url.Values{}

	for field, value := range websiteCredentials.LoginInfo {
		form.Add(field, value)
	}

	for field, value := range conf.Login.GenericFields {
		form.Add(field, value)
	}

	if len(conf.Login.ServerGeneratedFields) > 0 {
		serverGeneratedTokens := helpers.GetServerGeneratedTokens(conf.Login.Url, conf.Login.ServerGeneratedFields)
		for token, value := range serverGeneratedTokens {
			form.Add(token, value)
			fmt.Println(value)
		}
	}

	req, _ := http.NewRequest("POST", conf.Login.Url, strings.NewReader(form.Encode()))

	for header, value := range conf.Login.Headers {
		req.Header.Add(header, value)
	}

	if conf.Login.HomeUrl != "" {
		cookies := helpers.GetSiteCookies(conf.Login.HomeUrl)
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:109.0) Gecko/20100101 Firefox/109.0")

	// 	for item, value := range req.Header {
	// 		fmt.Println(item, " : ", value)
	// 	}
	// 	for item, value := range form {
	// 		fmt.Println(item, " : ", value)
	// 	}
	// for item, value := range req.Cookies() {
	// 		fmt.Println(item, " : ", value)
	// 	}
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

	return true
}
