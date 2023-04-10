package specificScrapers

import (
	"hatt/assets"
	"hatt/variables"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

func (t T) Memoryoftheworld() []variables.Item {
	config := assets.DeserializeWebsiteConf("memoryoftheworld.json")
	specificInfo := config.SpecificInfo

	apiUrl := specificInfo["apiUrl"]

	client := &http.Client{}
	formattedInput := strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement)
	requestUrl := apiUrl + formattedInput
	req, _ := http.NewRequest("GET", requestUrl, nil)

	response, _ := client.Do(req)

	responseData, _ := ioutil.ReadAll(response.Body)
	jsonResponse := string(responseData)

	searchResults := gjson.Get(jsonResponse, "_items").Array()

	var results []variables.Item
	for _, result := range searchResults {
		treatableResult := result.Raw
		var item variables.Item

		item.Link = specificInfo["itemBaseUrl"] + gjson.Get(treatableResult, "_id").Str
		item.Name = gjson.Get(treatableResult, "title").Str
		item.Thumbnail = "https:" + gjson.Get(treatableResult, "library_url").Str + gjson.Get(treatableResult, "cover_url").Str

		var authors string
		for index, author := range gjson.Get(treatableResult, "authors").Array() {
			if index != 0 {
				authors += ", "
			}
			authors += author.Str
		}
		item.Metadata = map[string]string{
			"autors": authors,
		}
		results = append(results, item)
	}

	return results

}
