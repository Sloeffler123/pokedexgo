package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://pokeapi.co/api/v2/location-area/"
const url2 = "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20"

type NamedResource struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type Location struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []NamedResource `json:"results"`
}

func findLoc() {
	var loc Location
	res, err := http.Get(url)
	if err != nil {
		fmt.Errorf("http get fail")
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("readall fail")
	}
	if err := json.Unmarshal(data, &loc); err != nil {
		fmt.Errorf("unmarshall did not work")
	}
	for _, value := range loc.Results {
		fmt.Println(value.Name)
	}
}
