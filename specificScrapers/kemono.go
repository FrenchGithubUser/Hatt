package specificScrapers

import (
	"hatt/assets"
	"hatt/variables"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/gocolly/colly"
	"github.com/tidwall/gjson"
)

func (t T) Kemono() []variables.Item {
	config := assets.DeserializeWebsiteConf("kemono.json")
	specificInfo := config.SpecificInfo

	apiUrl := specificInfo["apiUrl"]

	var formattedInput string = strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl, nil)

	response, _ := client.Do(req)

	responseData, _ := ioutil.ReadAll(response.Body)
	jsonResponse := string(responseData)

	creators := gjson.Parse(jsonResponse).Array()

	var results []variables.Item
	var wg sync.WaitGroup
	for _, creator := range creators {
		treatableResult := creator.Raw
		var item variables.Item
		item.Name = gjson.Get(treatableResult, "name").Str

		if strings.Contains(item.Name, formattedInput) {
			wg.Add(1)

			go func(result gjson.Result) {
				item.Link = specificInfo["itemBaseUrl"] + gjson.Get(treatableResult, "service").Str + "/user/" + gjson.Get(treatableResult, "id").Str
				item.Thumbnail = specificInfo["itemBaseUrl"] + "icons/" + gjson.Get(treatableResult, "service").Str + "/" + gjson.Get(treatableResult, "id").Str

				c := colly.NewCollector()
				c.OnHTML("#paginator-top", func(h *colly.HTMLElement) {
					pagination := strings.Split(h.ChildText("small"), "of ")
					if len(pagination) >= 2 {
						item.Metadata = map[string]string{
							"files": pagination[1] + " files",
						}
					}
				})
				c.Visit(item.Link)

				results = append(results, item)
				wg.Done()

			}(creator)
		}
	}
	wg.Wait()

	return results

}
