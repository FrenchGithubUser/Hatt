package specificScrapers

import (
	"encoding/json"
	"hatt/assets"
	"hatt/variables"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

type hotleakApiResponse struct {
	Models hotleakModelsInfo `json:"models"`
}

type hotleakModelsInfo struct {
	Data []hotleakModel `json:"data"`
}

type hotleakModel struct {
	Thumbnail string `json:"origin_image"`
	Name      string `json:"key"`
}

func (t T) Hotleak() []variables.Item {
	config := assets.DeserializeWebsiteConf("hotleak.json")
	specificInfo := config.SpecificInfo

	apiUrl := specificInfo["apiUrl"] + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl, nil)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	response, _ := client.Do(req)

	var jsonResponse hotleakApiResponse

	responseData, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(responseData, &jsonResponse)

	var results []variables.Item
	for _, value := range jsonResponse.Models.Data {
		var item variables.Item
		item.Link = "https://hotleak.vip/" + value.Name
		item.Name = value.Name
		//thumbnail requires cookie to be displayed, get it with the helper function
		// item.Thumbnail = "https://hotleak.vip/" + value.Thumbnail

		c := colly.NewCollector()

		c.OnHTML("section ul li a#all-tab", func(h *colly.HTMLElement) {
			item.Metadata = map[string]string{
				"files": strings.TrimSuffix(strings.Split(h.Text, "All (")[1], ")") + " files",
			}
		})

		c.Visit(item.Link)

		results = append(results, item)
	}

	return results

}
