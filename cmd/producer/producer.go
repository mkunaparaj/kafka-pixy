package main

import (
	"log"
	"time"

	"sandbox/kafka-pixy/clients/grpc/client"
)

const (
	addr     string = "localhost:19091"
	topic    string = "main_topic"
	interval int    = 5 // in seconds
)

func main() {

	c := client.New(addr)

	t := time.Tick(time.Duration(interval) * time.Second)

	for _ = range t {

		resp, err := c.Produce(topic, []byte("test_key"), []byte("test_msg"))
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("produced msg success: partition=%d, offset=%d", resp.Partition, resp.Offset)

	}
}
