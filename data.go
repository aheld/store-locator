package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
)

type Market struct {
	Id          string `json:"listing_id"`
	Name        string `json:"listing_name"`
	Address     string `json:"location_address"`
	Description string `json:"brief_desc"`
	Website     string `json:"media_website"`
	Latitude    string `json:"location_x"`
	Longitude   string `json:"location_y"`
	Image       string `json:"listing_image"`
	Products    []string
}

//go:embed data/farmers_markets_pa.json
var f embed.FS

func Import() ([]Market, error) {
	jsonFile, err := f.ReadFile("data/farmers_markets_pa.json")
	if err != nil {
		return nil, err
	}

	var markets []Market
	err = json.Unmarshal(jsonFile, &markets)
	for i := range markets {
		markets[i].Description, markets[i].Products = splitDescription(markets[i].Description)
	}
	return markets, err
}

func splitDescription(desc string) (string, []string) {
	description := desc
	var products []string = []string{}
	if strings.Contains(desc, "Products:") {

		tmp := strings.Split(desc, "Products: ")
		description = tmp[0]
		products = strings.Split(tmp[1], ";")
	}
	fmt.Println(description)
	description = strings.Replace(description, "<br>Available", "", 1)
	return description, products
}
