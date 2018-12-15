package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/anaskhan96/soup"
)

var products []Product

func parseProduct(result soup.Root, wg *sync.WaitGroup, mut *sync.Mutex) {
	product := Product{}

	product.Link = result.Find("a", "class", "s-access-detail-page").Attrs()["href"]
	product.Name = result.Find("h2", "class", "s-access-title").Text()
	product.Image = result.Find("img", "class", "s-access-image").Attrs()["src"]
	product.Price = result.Find("span", "class", "s-price").Text()

	product.GetReviews()

	mut.Lock()
	products = append(products, product)
	mut.Unlock()

	wg.Done()
}

func main() {
	now := time.Now().UTC()

	resp, err := soup.Get("https://www.amazon.in/TVs/b/ref=nav_shopall_sbc_tvelec_television?ie=UTF8&node=1389396031")

	// fmt.Println("Main fetch time: ", time.Since(now))
	// now = time.Now().UTC()

	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)
	results := doc.Find("div", "id", "mainResults").FindAll("li", "class", "s-result-item")

	var wg sync.WaitGroup
	var mut sync.Mutex
	for _, result := range results {
		wg.Add(1)
		go parseProduct(result, &wg, &mut)
	}

	wg.Wait()

	json.NewEncoder(os.Stdout).Encode(products)

	fmt.Printf("{\"time\": \"%s\", \"count\": %d}\n", time.Since(now), len(results))
}
