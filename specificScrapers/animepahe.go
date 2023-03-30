package specificScrapers

import (
	"hatt/assets"
	"hatt/variables"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

func (t T) Animepahe() []variables.Item {
	config := assets.DeserializeWebsiteConf("animepahe.json")
	specificInfo := config.SpecificInfo

	apiUrl := specificInfo["apiUrl"]

	client := &http.Client{}
	formattedInput := strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement)
	requestUrl := apiUrl + formattedInput
	req, _ := http.NewRequest("GET", requestUrl, nil)

	response, _ := client.Do(req)

	responseData, _ := ioutil.ReadAll(response.Body)
	jsonResponse := string(responseData)

	searchResults := gjson.Get(jsonResponse, "data").Array()

	var results []variables.Item
	for _, result := range searchResults {
		treatableResult := result.Raw
		var item variables.Item

		item.Name = gjson.Get(treatableResult, "title").Str
		item.Link = specificInfo["baseUrl"] + "anime/" + gjson.Get(treatableResult, "session").Str
		item.Thumbnail = gjson.Get(treatableResult, "poster").Str
		item.Metadata = map[string]string{
			"date":     gjson.Get(treatableResult, "year").Raw,
			"status":   gjson.Get(treatableResult, "status").Str,
			"type":     gjson.Get(treatableResult, "type").Str,
			"episodes": gjson.Get(treatableResult, "episodes").Raw + " episodes",
		}

		results = append(results, item)
	}

	return results

}
