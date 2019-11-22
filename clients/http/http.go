package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	baseAddr    string = "http://localhost:19092"
	producePath string = "/topics/%s/messages"
	consumePath string = "/topics/%s/messages?group=%s"

	baseMsg   string = `{"Message": "Hello_%d", "Array": [1, 2, 3], "Null": null, "Number": 1.234}`
	topicName string = "main_topic"
	groupName string = "main_consumer_group"
)

func main() {

	//if err := produce(topicName, 10000); err != nil {
	//	log.Fatalf("produce error: %s", err.Error())
	//}

	if err := consume(topicName, groupName); err != nil {
		log.Fatalf("consume error: %s", err.Error())
	}
}

func consume(topic, group string) error {

	topicPath := fmt.Sprintf(consumePath, topic, group)

	// form requests
	req, err := formRequest(http.MethodGet, fmt.Sprintf("%s%s", baseAddr, topicPath), "")
	if err != nil {
		return err
	}

	// do http request
	if err := do(req); err != nil {
		return err
	}

	return nil
}

func produce(topic string, n int64) error {

	// form messages
	msgs := formMsgs(n)

	reqs := make([]*http.Request, len(msgs))

	// form http requests
	for i, m := range msgs {

		topicPath := fmt.Sprintf(producePath, topic)

		req, err := formRequest(http.MethodPost, fmt.Sprintf("%s%s", baseAddr, topicPath), m)
		if err != nil {
			return err
		}

		reqs[i] = req
	}

	// do http requests
	var wg sync.WaitGroup
	wg.Add(len(reqs))
	for _, r := range reqs {

		go func(r *http.Request) {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			defer wg.Done()
			if err := do(r); err != nil {
				log.Fatalln(err)
			}

		}(r)
	}
	wg.Wait()

	return nil
}

func do(req *http.Request) error {

	c := &http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		body, err := getBody(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	body, err := getBody(resp.Body)
	if err != nil {
		return err
	}

	log.Println(body)

	return nil
}

func getBody(body io.Reader) (string, error) {

	b, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func formRequest(method, url, body string) (*http.Request, error) {

	var req *http.Request
	var err error

	if body == "" {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, url, strings.NewReader(body))
		if err != nil {
			return nil, err
		}
	}

	// headers
	req.Header.Add("content-type", "application/json")

	return req, nil
}

func formMsgs(n int64) []string {

	msgs := make([]string, n)

	var i int64
	for i = 0; i < n; i++ {

		k := time.Now().UnixNano() + int64(i)
		m := fmt.Sprintf(baseMsg, k)
		msgs[i] = m
	}

	return msgs
}
