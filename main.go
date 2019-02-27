package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/anaskhan96/soup"
)

func parseProduct(result soup.Root) {
	product := Product{}

	product.Link = result.Find("a", "class", "s-access-detail-page").Attrs()["href"]
	product.Name = result.Find("h2", "class", "s-access-title").Text()
	product.Image = result.Find("img", "class", "s-access-image").Attrs()["src"]
	product.Price = result.Find("span", "class", "s-price").Text()

	product.GetReviews()

	json.NewEncoder(os.Stdout).Encode(product)
}

func main() {
	now := time.Now().UTC()

	resp, err := soup.Get("https://www.amazon.in/b/?_encoding=UTF8&node=1389401031&ref_=sv_top_elec_mega_1")

	// fmt.Println("Main fetch time: ", time.Since(now))
	// now = time.Now().UTC()

	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)
	results := doc.Find("div", "id", "mainResults").FindAll("li", "class", "s-result-item")

	for _, result := range results {
		parseProduct(result)
	}

	fmt.Printf("{\"time\": \"%s\", \"count\": %d}\n", time.Since(now), len(results))
}
