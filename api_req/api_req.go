package apireq

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Config struct {
	NextURL string
	PrevURL string
	ID      int
}

type LocationArea struct {
	ID   int
	Name string
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

func RequestToLocation(data []byte) (LocationArea, error) {
	location := LocationArea{}
	err := json.Unmarshal(data, &location)
	if err != nil {
		fmt.Println(err)
		return location, err
	}
	return location, nil
}
