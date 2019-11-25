package main

import (
	"context"
	"log"
	pb "sandbox/kafka-pixy/clients/grpc"
	"sandbox/kafka-pixy/clients/grpc/client"
)

const (
	addr    string = "localhost:19091"
	topic   string = "main_topic"
	group   string = "5"
	cluster string = "default"
)

func main() {

	c := client.New(addr)

	for {

		cons := &pb.ConsNAckRq{
			Cluster: cluster,
			Topic:   topic,
			Group:   group,
		}

		consRs, err := c.ConsumeNAck(context.TODO(), cons)
		if err != nil {
			log.Fatalln(err)
		}

		ack := &pb.AckRq{
			Cluster:   cluster,
			Topic:     topic,
			Group:     group,
			Partition: consRs.GetPartition(),
			Offset:    consRs.GetOffset(),
		}

		_, err = c.Ack(context.TODO(), ack)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("consumed msg: partition=%d offset=%d", consRs.GetPartition(), consRs.GetOffset())
	}
}
