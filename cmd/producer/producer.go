package main

import (
	"context"
	"log"
	pb "sandbox/kafka-pixy/clients/grpc"
	"sandbox/kafka-pixy/clients/grpc/client"
)

const (
	addr     string = "localhost:19091"
	topic    string = "main_topic"
	group    string = "consumer_group_5"
	cluster  string = "default"
	interval int    = 1 // in seconds
)

func main() {

	c := client.New(addr)

	for {

		input := &pb.ProdRq{
			Cluster:  cluster,
			Topic:    topic,
			KeyValue: []byte("test_key"),
			Message:  []byte("test_msg"),
		}

		rs, err := c.Produce(context.TODO(), input)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("produced msg: partition=%d offset=%d", rs.GetPartition(), rs.GetOffset())
	}
}
