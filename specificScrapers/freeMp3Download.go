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

func (t T) FreeMp3Download() []variables.Item {
	config := assets.DeserializeWebsiteConf("freeMp3Download.json")
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

		// this url is the direct link to the song's detail page, but it needs the "Referer" header to be displayed, which isn't the case when copy/pasted
		// item.Link = specificInfo["detailUrl"] + "?id=" + gjson.Get(treatableResult, "id").Str + "&q=" + base64.StdEncoding.EncodeToString([]byte(formattedInput))
		item.Link = specificInfo["baseUrl"]

		item.Name = gjson.Get(treatableResult, "title").Str
		item.Thumbnail = gjson.Get(treatableResult, "album.cover_medium").Str
		item.Metadata = map[string]string{
			"duration": helpers.FormatDuration(int(gjson.Get(treatableResult, "duration").Int())),
			"artist":   gjson.Get(treatableResult, "artist.name").Str,
		}
		results = append(results, item)
	}

	return results

}
