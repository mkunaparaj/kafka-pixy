package client

import (
	"log"
	pb "github.com/sendgrid/kafka-pixy-POC/clients/grpc"

	"google.golang.org/grpc"
)

func New(addr string) pb.KafkaPixyClient {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	return pb.NewKafkaPixyClient(conn)
}
