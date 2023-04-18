package login

import (
	"fmt"
	"hatt/assets"
	"hatt/configuration"
	"hatt/helpers"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func LoginBrowser(website string) bool /*return whether or not the login was successfull*/ {

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
		return true
	}

	websiteCredentials.Tokens = map[string]map[string]string{}

	l := helpers.InstanciateBrowser()
	page := rod.New().ControlURL(l).MustConnect().MustPage(conf.Login.Url)
	page.MustWaitLoad()

	for field, value := range websiteCredentials.LoginInfo {
		page.MustElement(conf.Login.PageInputs[field]).MustClick().MustInput(value)
	}

	// clicks on the login button and waits for the response from the server
	var wg sync.WaitGroup
	wg.Add(1)
	go page.EachEvent(func(e *proto.PageLoadEventFired) {
		// page loaded
		wg.Done()
	})()
	page.MustElement(conf.Login.LoginButton).MustClick()
	// page.MustScreenshot("a.png")
	wg.Wait()

	cookies, _ := page.Browser().GetCookies()
	for _, cookie := range cookies {
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
