package client

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

type Client interface {
	Produce(topic string, key, msg []byte) (*pb.ProdRs, error)
	Consume(topic, group string) (*pb.ConsRs, error)
}

type client struct {
	c pb.KafkaPixyClient
}

func New(addr string) Client {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	k := pb.NewKafkaPixyClient(conn)

	return &client{c: k}
}

func (c *client) Produce(topic string, key, msg []byte) (*pb.ProdRs, error) {

	req := &pb.ProdRq{
		Topic:    topic,
		KeyValue: key,
		Message:  msg,
	}

	resp, err := c.c.Produce(context.TODO(), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) Consume(topic, group string) (*pb.ConsRs, error) {

	req := &pb.ConsNAckRq{
		Topic: topic,
		Group: group,
	}

	resp, err := c.c.ConsumeNAck(context.TODO(), req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
