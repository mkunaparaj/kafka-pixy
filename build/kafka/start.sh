#!/bin/bash

echo 'starting zookeeper with configuration:'
cat ./kafka/config/zookeeper.properties | echo
./kafka/bin/zookeeper-server-start.sh ./kafka/config/zookeeper.properties > zk-output.log &

echo 'starting kafka with configuration:'
cat ./kafka/config/server.properties | echo
./kafka/bin/kafka-server-start.sh ./kafka/config/server.properties
