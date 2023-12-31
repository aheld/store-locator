package main

import (
	"embed"
	"encoding/json"
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
	return markets, err
}
