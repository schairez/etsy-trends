package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	etsyAPIKey := os.Getenv("ETSY_API_KEY")
	etsyAPIURL := getActiveListings("face+mask", etsyAPIKey)

	fmt.Println(etsyAPIURL)
	resp, err := http.Get(etsyAPIURL)
	if err != nil {
		log.Fatalln(err)

	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var ResponseStruct APIResponse
	json.Unmarshal(bodyBytes, &ResponseStruct)

	// fmt.Printf("API Response as struct %+v\n", ResponseStruct)
	//filter by most views, or most sales?
	//figure out trending keywords / tags for search listings
	//sort on views?
	//https://openapi.etsy.com/v2/listings/active?keywords=face+mask&sort_on=score&api_key=i1l7s0q019amb0m4nqmv267r

	ResponseJSON, _ := json.MarshalIndent(ResponseStruct, "", "\t")
	err = ioutil.WriteFile("output.json", ResponseJSON, 0644)

}

func getEtsyBaseFeed() string {
	return "https://openapi.etsy.com/v2/feeds"
}

func getActiveListings(keywords, etsyAPIKey string) string {
	return fmt.Sprintf("https://openapi.etsy.com/v2/listings/active?sort_on=score&keywords=%sw&api_key=%s", keywords, etsyAPIKey)
}

func getTrendingListings(etsyAPIKey string) string {
	return fmt.Sprintf("https://openapi.etsy.com/v2/listings/trending?&api_key=%s", etsyAPIKey)

}

// APIResponse struct
type APIResponse struct {
	Count   int      `json:"count"`
	Results []Result `json:"results"`
}

// Result struct
type Result struct {
	ListingID    int      `json:"listing_id"`
	State        string   `json:"state"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Price        string   `json:"price"`
	Quantity     int      `json:"quantity"`
	Tags         []string `json:"tags"`
	Materials    []string `json:"materials"`
	Views        int      `json:"views"`
	TaxonomyPath []string `json:"taxonomy_path"`
}
