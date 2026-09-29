package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

type MapResponse struct {
	Count		int
	Next 		string
	Previous 	string
	Results		[]struct {
		Name 	string
		Url 	string
	}
}

func commandMap(cfg *config) error {
	res, err := http.Get(cfg.nextURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var apiResponse MapResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}
	for _,location := range apiResponse.Results {
		fmt.Println(location.Name)
	}
	cfg.prevURL = apiResponse.Previous
	cfg.nextURL = apiResponse.Next
	return nil
}

func commandMapb(cfg *config) error {
	res, err := http.Get(cfg.prevURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var apiResponse MapResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}
	for _,location := range apiResponse.Results {
		fmt.Println(location.Name)
	}
	cfg.prevURL = apiResponse.Previous
	cfg.nextURL = apiResponse.Next
	return nil
}