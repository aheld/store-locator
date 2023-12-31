package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data/farmers_markets_from_usda.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading records")
	}

	if len(records) < 2 {
		log.Fatal("File must have 2 or more rows")
	}

	headers := records[0]
	fmt.Println(headers)
	entries := []map[string]string{}
	for _, eachrecord := range records[1:] {
		entry := make(map[string]string)
		for i, value := range eachrecord {
			entry[headers[i]] = value
		}
		entries = append(entries, entry)
	}
	bytes, err := json.MarshalIndent(entries, "", "	")
	if err != nil {
		log.Fatalf("Marshal error %s\n", err)
	}

	fmt.Println(string(bytes))
}
