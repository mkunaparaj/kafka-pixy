package main

import (
	"context"
	"log"
	pb "sandbox/kafka-pixy/clients/grpc"

	"google.golang.org/grpc"
)

const (
	addr      string = "localhost:19091"
	topicName string = "main_topic"
)

func main() {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := pb.NewKafkaPixyClient(conn)

	request := &pb.ProdRq{
		Topic:    topicName,
		KeyValue: []byte("test_key"),
		Message:  []byte("test_msg"),
	}
	resp, err := client.Produce(context.TODO(), request)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("produced msg success: partition=%d, offset=%d", resp.Partition, resp.Offset)
}
