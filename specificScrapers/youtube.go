package specificScrapers

import (
	"hatt/assets"
	"hatt/variables"
	"strings"

	"github.com/tidwall/gjson"

	"github.com/gocolly/colly"
)

func (t T) Youtube() []variables.Item {

	var results []variables.Item
	c := colly.NewCollector()

	config := assets.DeserializeWebsiteConf("youtube.json")

	c.OnHTML("body > script:nth-of-type(14)", func(h *colly.HTMLElement) {

		jsonData := h.Text[strings.Index(h.Text, "{") : len(h.Text)-1]
		contentsList := gjson.Get(jsonData, "contents.twoColumnSearchResultsRenderer.primaryContents.sectionListRenderer.contents").Array()[0]
		ytResultsList := gjson.Get(contentsList.Raw, "itemSectionRenderer.contents").Array()

		for _, ytResult := range ytResultsList {
			result := gjson.Get(ytResult.Raw, "videoRenderer").Raw

			// some videos don't have a title for some reason
			if gjson.Get(result, "title.runs").Raw != "" {
				item := variables.Item{
					Name:      gjson.Get(gjson.Get(result, "title.runs").Array()[0].Raw, "text").Str,
					Link:      "https://youtube.com/watch?v=" + gjson.Get(result, "videoId").Str,
					Thumbnail: gjson.Get(gjson.Get(result, "thumbnail.thumbnails").Array()[0].Raw, "url").Str,
				}

				item.Metadata = map[string]string{
					"published_at": gjson.Get(result, "publishedTimeText.simpleText").Str,
					"views":        gjson.Get(result, "shortViewCountText.simpleText").Str,
					"duration":     gjson.Get(result, "lengthText.simpleText").Str,
					"author":       gjson.Get(gjson.Get(result, "longBylineText.runs").Array()[0].Raw, "text").Str,
				}
				results = append(results, item)
			}

		}

	})

	c.Visit(config.Search.Url + strings.ReplaceAll(variables.CURRENT_INPUT, " ", config.Search.SpaceReplacement))

	return results
}
