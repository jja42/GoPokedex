package apireq

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Config struct {
	NextURL *string
	PrevURL *string
}

type RespLocationBatch struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetRequest(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d\n", res.StatusCode)
	}
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func RequestToLocations(data []byte) (RespLocationBatch, error) {
	locations := RespLocationBatch{}
	err := json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Println(err)
		return locations, err
	}
	return locations, nil
}
