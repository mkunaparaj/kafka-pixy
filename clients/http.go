package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseAddr string = "localhost:19092"
)

func main() {

	c := &http.Client{}

	req, err := formRequest("GET", "/topics")
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("response body: %s\n", body)
}

func formRequest(method, path string) (*http.Request, error) {

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", baseAddr, path), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
