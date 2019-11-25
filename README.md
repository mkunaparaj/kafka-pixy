# kafka-pixy-POC

## Startup order

1. Zookeeper
2. Kafka
3. Pixy
4. UI

## TODO

* form kafka cluster not just single node
* topic with multiple paritions then use consumer group on it
* retries on LeaderNotAvailable errors
* Kafka consistency patterns
* own grpc protocol
* multiple kafka clusters
* seem to be losing messages in between start up and down

## Curls

```bash
curl localhost:19092/topics
curl 'http://localhost:19092/topics/topic2/messages?group=1'
```

## Kafkacat

run from jump

```bash
kubectl get pods | grep jump
kubectl exec -it jumpbox bash

kafkacat -b 10.96.1.1:9092 -L
kafkacat -b 10.96.1.1:9092 -t topic1 -P
```
