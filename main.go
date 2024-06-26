// A web scraper that will take an address,
// (eg. www.google.com)
// scrape it for an html element,
// then display it.
package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/mikevidotto/greeting"
	"strconv"
	"strings"
)

func main() {
	greet := greetings.Hello("Michael")
	fmt.Println(greet)

	ScrapeUrl := "https://github.com/mikevidotto"

	c := colly.NewCollector(colly.AllowedDomains("www.github.com", "github.com"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String(), "\n")
	})

	c.OnHTML("h1.vcard-names span.p-nickname", func(h *colly.HTMLElement) {
		fmt.Println(strings.TrimSpace(h.Text) + "'s repos\n")
	})

	c.OnHTML("li.mb-3 span.repo", func(h *colly.HTMLElement) {
		fmt.Println("repo " + strconv.Itoa(h.Index+1) + ": " + strings.TrimSpace(h.Text))
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error encountered: ", err.Error())
	})

	c.Visit(ScrapeUrl)
}
