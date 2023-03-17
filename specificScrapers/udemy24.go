package specificScrapers

import (
	"hatt/assets"
	"hatt/variables"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

func (t T) Udemy24() []variables.Item {
	config := assets.DeserializeWebsiteConf("udemy24.json")
	specificInfo := config.SpecificInfo

	apiUrl := specificInfo["apiUrl"]

	client := &http.Client{}
	formattedInput := strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement)
	requestUrl := apiUrl + formattedInput
	req, _ := http.NewRequest("GET", requestUrl, nil)

	response, _ := client.Do(req)

	responseData, _ := ioutil.ReadAll(response.Body)
	jsonResponse := string(responseData)

	searchResults := gjson.Get(jsonResponse, "results").Array()

	var results []variables.Item
	for _, result := range searchResults {
		treatableResult := result.Raw
		var item variables.Item

		item.Name = gjson.Get(treatableResult, "fields.title\\.default").Str
		item.Link = "https://" + gjson.Get(treatableResult, "fields.permalink\\.url\\.raw").Str
		item.Thumbnail = "https://" + gjson.Get(treatableResult, "fields.image\\.url\\.raw").Str
		item.Metadata = map[string]string{
			"date": gjson.Get(treatableResult, "fields.date").Str,
		}

		results = append(results, item)
	}

	return results

}
