package scrap

import (
	"github.com/gocolly/colly"
)

type Scrap interface {
	Scrap_links() string
}

type Scrap_fields struct {
	Web_address    string
	Allowed_domain string
}

//Find all links to a website
func (field Scrap_fields) Scrap_links() string {
	var str string
	c := colly.NewCollector(colly.AllowedDomains(field.Allowed_domain))

	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		for _, value := range links {
			str = str + " " + value + " " + "\n"
		}
	})

	c.Visit(field.Web_address)
	return str
}
