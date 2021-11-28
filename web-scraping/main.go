package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.globalbrand.com.bd/product-category/components/monitor-price-in-bangladesh/"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var names []string
	var prices []string
	var links []string

	document.Find(".product").Each(func(i int, s *goquery.Selection) {
		// names
		name := s.Find("h2").Text()
		names = append(names, name)

		// links
		link, exist := s.Find("a").Attr("href")
		if !exist {
			link = "no url found"
		}
		links = append(links, link)

		//prices
		price := s.Find(".price").Text()
		if price == "" {
			price = "UPCOMING !!!"
		}
		prices = append(prices, price)
	})

	for key := range names {
		fmt.Println("Name:  ", names[key])
		fmt.Println("Price: ", prices[key])
		fmt.Println("Link:  ", links[key])
	}
}
