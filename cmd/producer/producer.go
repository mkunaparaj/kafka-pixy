package main

import (
	"context"
	"fmt"
	"log"
	pb "github.com/sendgrid/kafka-pixy-POC/clients/grpc"
	"github.com/sendgrid/kafka-pixy-POC/clients/grpc/client"
)

const (
	addr     string = "localhost:19091"
	topic    string = "main_topic"
	group    string = "4"
	cluster  string = "default"
	interval int    = 1 // in seconds
)

func main() {

	c := client.New(addr)

	var count int
	for {

		input := &pb.ProdRq{
			Cluster:  cluster,
			Topic:    topic,
			KeyValue: []byte(fmt.Sprintf("test_key_%d", count)),
			Message:  []byte("test_msg"),
			//KeyUndefined: true,
		}

		rs, err := c.Produce(context.TODO(), input)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("produced msg: partition=%d offset=%d", rs.GetPartition(), rs.GetOffset())

		count++
	}
}
