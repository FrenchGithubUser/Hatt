package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	Name     string
	ThumbUrl string
}

func main() {
	c := colly.NewCollector()

	var items []item

	// item information
	c.OnHTML(".col-sm", func(h *colly.HTMLElement) {
		// fmt.Println(h.ChildText("h2"))
		item := item{
			Name:     h.ChildText("div.file-right a.ai-search h2"),
			ThumbUrl: h.ChildAttr("div.file-left a img.img-zoom", "src"),
		}

		items = append(items, item)
		fmt.Println(item)
	})

	// pagination handling
	c.OnHTML("a.navigation.next", func(h *colly.HTMLElement) {
		nextPage := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(nextPage)

	})

	c.Visit("https://www.pdfdrive.com/search?q=test&pagecount=&pubyear=&searchin=&em=")
	// fmt.Println(items)
}
