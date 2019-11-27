package main

import (
	"context"
	"log"
	pb "github.com/sendgrid/kafka-pixy/clients/grpc"
	"github.com/sendgrid/kafka-pixy/clients/grpc/client"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	addr    string = "localhost:19091"
	topic   string = "main_topic"
	group   string = "20"
	cluster string = "default"
)

func main() {

	c := client.New(addr)

	cons := &pb.ConsNAckRq{
		Cluster: cluster,
		Topic:   topic,
		Group:   group,
		NoAck:   true,
	}

	consRs, err := c.ConsumeNAck(context.TODO(), cons)
	if err != nil {
		log.Fatalln(err)
	}

	ackPartition := consRs.GetPartition()
	ackOffset := consRs.GetOffset()
	for {

		cons := &pb.ConsNAckRq{
			Cluster:      cluster,
			Topic:        topic,
			Group:        group,
			AckPartition: ackPartition, // ack previous message
			AckOffset:    ackOffset,    // ack previous message
		}

		consRs, err := c.ConsumeNAck(context.TODO(), cons)
		if err != nil {

			// check if no messages in bus
			// if no messages, exit for loop and ack last message then exit app
			if status.Code(err) == codes.NotFound {
				log.Println(err)
				break
			} else {
				log.Fatalln(err)
			}
		}

		ackPartition = consRs.GetPartition()
		ackOffset = consRs.GetOffset()
		log.Printf("consumed msg: partition=%d offset=%d", consRs.GetPartition(), consRs.GetOffset())
	}

	// ack last message
	last := &pb.AckRq{
		Cluster:   cluster,
		Topic:     topic,
		Group:     group,
		Partition: consRs.GetPartition(),
		Offset:    consRs.GetOffset(),
	}

	_, err = c.Ack(context.TODO(), last)
	if err != nil {
		log.Println("HERE")
		log.Fatalln(err)
	}

	log.Printf("finished, exiting")
}
