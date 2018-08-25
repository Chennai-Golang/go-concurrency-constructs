package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/anaskhan96/soup"
)

func parseProducts(result soup.Root) Product {
	product := Product{}

	product.Link = result.Find("a", "class", "s-access-detail-page").Attrs()["href"]
	product.Name = result.Find("h2", "class", "s-access-title").Text()
	product.Image = result.Find("img", "class", "s-access-image").Attrs()["src"]
	product.Price = result.Find("span", "class", "s-price").Text()

	product.GetReviews()

	return product
}

func main() {
	now := time.Now().UTC()

	resp, err := soup.Get("https://www.amazon.in/TVs/b/ref=nav_shopall_sbc_tvelec_television?ie=UTF8&node=1389396031")

	fmt.Println("Main fetch time: ", time.Since(now))
	now = time.Now().UTC()

	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)
	results := doc.Find("div", "id", "mainResults").FindAll("li", "class", "s-result-item")

	for _, result := range results {
		json.NewEncoder(os.Stdout).Encode(parseProducts(result))
	}

	fmt.Println("Elapsed time: ", time.Since(now))
}
