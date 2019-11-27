package main

import (
	"context"
	"log"
	pb "github.com/sendgrid/kafka-pixy-POC/clients/grpc"
	"github.com/sendgrid/kafka-pixy-POC/clients/grpc/client"
)

const (
	addr    string = "localhost:19091"
	topic   string = "main_topic"
	group   string = "20"
	cluster string = "default"
)

func main() {

	c := client.New(addr)

	po := &pb.PartitionOffset{
		Partition: 0,
		Offset:    50,
	}

	p := make([]*pb.PartitionOffset, 0)

	p = append(p, po)

	input := &pb.SetOffsetsRq{
		Cluster: cluster,
		Topic:   topic,
		Group:   group,
		Offsets: p,
	}

	// does not effect existing consumer groups
	// to set offset on a currently running consumer group
	// it's necessary to delete the group or wait for the timeout in kafka-pixy
	// then run this command to set the offset to the desired number
	_, err := c.SetOffsets(context.TODO(), input)
	if err != nil {
		log.Fatalln(err)
	}
}
