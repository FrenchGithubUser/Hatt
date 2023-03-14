package specificScrapers

import (
	"encoding/json"
	"hatt/assets"
	"hatt/variables"
	"io/ioutil"
	"net/http"
	"strings"
)

type bilibiliApiResponse struct {
	Data data `json:"data"`
}

type data struct {
	Result []result `json:"result"`
}

type result struct {
	ResultType string      `json:"result_type"`
	Data       []videoInfo `json:"data"`
}

type videoInfo struct {
	VideoId   string `json:"bvid"`
	Name      string `json:"title"`
	Thumbnail string `json:"pic"`
	Duration  string `json:"duration"`
}

func (t T) Bilibili() []variables.Item {
	config := assets.DeserializeWebsiteConf("bilibili.json")
	specificInfo := config.SpecificInfo

	apiUrl := specificInfo["apiUrl"]

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl+strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement), nil)
	cookies := getBilibiliCookies()
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:109.0) Gecko/20100101 Firefox/109.0")
	response, _ := client.Do(req)

	var jsonResponse bilibiliApiResponse

	responseData, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(responseData, &jsonResponse)

	var results []variables.Item
	for _, value := range jsonResponse.Data.Result {
		if value.ResultType == "video" {
			for _, video := range value.Data {
				var item variables.Item
				var videoInfo videoInfo = video
				item.Link = "https://www.bilibili.com/video/" + videoInfo.VideoId
				item.Name = videoInfo.Name
				item.Thumbnail = "https://" + videoInfo.Thumbnail[2:]
				item.Metadata = map[string]string{
					"duration": videoInfo.Duration,
				}
				results = append(results, item)
			}
		}
	}

	return results

}

// cookies are needed to set api requests (cookies generated when visiting the home page)
func getBilibiliCookies() []*http.Cookie {
	response, _ := http.Get("https://www.bilibili.com/")
	return response.Cookies()
}
