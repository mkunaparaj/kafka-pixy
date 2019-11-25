package main

import (
	"log"
	"sandbox/kafka-pixy/clients/grpc/client"
	"time"
)

const (
	addr     string = "localhost:19091"
	topic    string = "main_topic"
	group    string = "consumer_group_1"
	interval int    = 2 // in seconds
)

func main() {

	c := client.New(addr)

	t := time.Tick(time.Duration(interval) * time.Second)

	for _ = range t {

		resp, err := c.Consume(topic, group)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("consumed msg success: key=%s partition=%d, offset=%d", resp.GetKeyValue(), resp.GetPartition(), resp.GetOffset())
	}
}
