package specificScrapers

import (
	"hatt/assets"
	"hatt/helpers"
	"hatt/variables"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

func (t T) Slavart() []variables.Item {
	config := assets.DeserializeWebsiteConf("slavart.json")
	specificInfo := config.SpecificInfo

	apiUrl := specificInfo["apiUrl"]

	client := &http.Client{}
	formattedInput := strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement)
	requestUrl := apiUrl + formattedInput
	req, _ := http.NewRequest("GET", requestUrl, nil)

	response, _ := client.Do(req)

	responseData, _ := ioutil.ReadAll(response.Body)
	jsonResponse := string(responseData)

	searchResults := gjson.Get(jsonResponse, "tracks.items").Array()

	var results []variables.Item
	for _, result := range searchResults {
		treatableResult := result.Raw
		var item variables.Item

		item.Link = specificInfo["baseUrl"]
		item.Name = gjson.Get(treatableResult, "title").Str
		item.Thumbnail = gjson.Get(treatableResult, "album.image.small").Str
		item.Metadata = map[string]string{
			"duration": helpers.FormatDuration(int(gjson.Get(treatableResult, "duration").Int())),
			"artist":   gjson.Get(treatableResult, "performer.name").Str,
		}
		results = append(results, item)
	}

	return results

}
