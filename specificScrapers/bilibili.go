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
	Data Data `json:"data"`
}

type Data struct {
	Result []Result `json:"result"`
}

type Result struct {
	ResultType string      `json:"result_type"`
	Data       []VideoInfo `json:"data"`
}

type VideoInfo struct {
	VideoId   string `json:"bvid"`
	Name      string `json:"title"`
	Thumbnail string `json:"pic"`
}

func (t T) Bilibili() []variables.Item {
	config := assets.DeserializeWebsiteConf("bilibili.json")
	specificInfo := config.SpecificInfo

	apiUrl := specificInfo["apiUrl"]

	response, _ := http.Get(apiUrl + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement))

	var jsonResponse bilibiliApiResponse

	responseData, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(responseData, &jsonResponse)

	var results []variables.Item
	for _, value := range jsonResponse.Data.Result {
		if value.ResultType == "video" {
			for _, video := range value.Data {
				var item variables.Item
				var videoInfo VideoInfo = video
				item.Link = "https://www.bilibili.com/video/" + videoInfo.VideoId
				item.Name = videoInfo.Name
				item.Thumbnail = "https://" + videoInfo.Thumbnail[2:]
				results = append(results, item)
			}
		}
	}

	return results

}
