#!/bin/bash

mkdir kafka-pixy
mv kafkapixy.proto kafka-pixy
protoc -I kafka-pixy/ kafkapixy.proto --go_out=plugins=grpc:kafka-pixy
