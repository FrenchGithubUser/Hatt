package specificScrapers

import (
	"encoding/json"
	"hatt/assets"
	"hatt/variables"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

type coomerCreator struct {
	Service string `json:"service"`
	Name    string `json:"name"`
}

func (t T) Coomer() []variables.Item {
	config := assets.DeserializeWebsiteConf("coomer.json")
	specificInfo := config.SpecificInfo

	apiUrl := specificInfo["apiUrl"]

	response, _ := http.Get(apiUrl)

	var jsonResponse []coomerCreator

	responseData, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(responseData, &jsonResponse)

	var results []variables.Item
	var formattedInput string = strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement)

	var wg sync.WaitGroup
	for _, value := range jsonResponse {
		if strings.Contains(value.Name, formattedInput) {

			wg.Add(1)
			go func(value coomerCreator) {
				var item variables.Item
				item.Link = "https://coomer.party/" + value.Service + "/user/" + value.Name
				item.Name = value.Name
				item.Thumbnail = "https://coomer.party/thumbnail/icons/" + value.Service + "/" + value.Name

				c := colly.NewCollector()

				c.OnHTML("#paginator-bottom", func(h *colly.HTMLElement) {
					if h.ChildText("small") != "" {
						item.Metadata = map[string]string{
							"files": strings.Split(h.ChildText("small"), "of ")[1] + " files",
						}
					}
				})

				c.Visit(item.Link)

				results = append(results, item)
				wg.Done()

			}(value)

		}
	}
	wg.Wait()

	return results

}
