package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Bytes []byte
var ShortBytes []byte

type City struct {
	ID        int
	Name      string
	Latitude  string
	Longitude string
}

type State struct {
	ID        int
	Name      string
	StateCode string `json:"state_code"`
	Latitude  string
	Longitude string
	CountryID int `json:"country_id"`
	Cities    []City
}

func readFile(name string) []byte {
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return bytes
}

func init() {
	Bytes = readFile("states-cities.json")
	ShortBytes = readFile("states-cities-short.json")
}

func main() {
	// For the sake of example...
	result := []State{}
	err := json.Unmarshal(Bytes, &result)
	if err != nil {
		panic(err)
	}
	cities := 0
	for _, state := range result {
		cities += len(state.Cities)
	}
	fmt.Printf("Processed %d states with %d nested cities\n", len(result), cities)
}
