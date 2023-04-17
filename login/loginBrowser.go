package login

import (
	"fmt"
	"hatt/assets"
	"hatt/configuration"
	"hatt/helpers"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func LoginBrowser(website string) {

	var conf configuration.Config = assets.DeserializeWebsiteConf(website + ".json")
	var websiteCredentials helpers.WebsiteCredentials
	h := &helpers.Helper{}
	websiteCredentials = h.DeserializeCredentials(website)

	isLoginNeeded := helpers.IsLoginNeeded(websiteCredentials, conf)

	if !isLoginNeeded {
		return
	}

	l := helpers.InstanciateBrowser()
	page := rod.New().ControlURL(l).MustConnect().MustPage(conf.Login.Url)
	page.MustWaitLoad()

	for field, value := range websiteCredentials.LoginInfo {
		page.MustElement(conf.Login.PageInputs[field]).MustClick().MustInput(value)
	}

	evt := proto.NetworkResponseReceived{}
	wait := page.WaitEvent(&evt)
	page.MustElement(conf.Login.LoginButton).MustClick()
	wait()

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

}
