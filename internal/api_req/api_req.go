package apireq

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jja42/GoPokedex/internal/pokecache"
)

type Config struct {
	NextURL *string
	PrevURL *string
	Cache   *pokecache.Cache
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

type LocationArea struct {
	ID                 int
	Name               string
	Pokemon_Encounters []struct {
		Pokemon struct {
			Name string
		}
	}
}

func GetRequest(url string, config *Config) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d\n", res.StatusCode)
	}
	if err != nil {
		log.Fatal(err)
	}

	config.Cache.Add(url, data)

	return data
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

func RequestToLocation(data []byte) (LocationArea, error) {
	location := LocationArea{}
	err := json.Unmarshal(data, &location)
	if err != nil {
		fmt.Println(err)
		return location, err
	}
	return location, nil
}
