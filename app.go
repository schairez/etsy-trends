package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	etsyAPIKey := os.Getenv("ETSY_API_KEY")
	// etsyURI := "https://openapi.etsy.com/v2/listings/trending?active&title=facemask&api_key="
	etsyURI := "https://openapi.etsy.com/v2/listings/active?keywords=face+mask&api_key="
	etsyAPIURL := etsyURI + etsyAPIKey
	fmt.Println(etsyAPIURL)
	resp, err := http.Get(etsyAPIURL)
	if err != nil {
		log.Fatalln(err)

	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var ResponseStruct APIResponse
	json.Unmarshal(bodyBytes, &ResponseStruct)

	fmt.Printf("API Response as struct %+v\n", ResponseStruct)

	ResponseJSON, _ := json.MarshalIndent(ResponseStruct, "", "\t")
	err = ioutil.WriteFile("output.json", ResponseJSON, 0644)

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

// func get() {+
// 	// resp, err := http.Get("")

// }
