# kafka-pixy-POC

## TODO

* separate kafka and zk
* form kafka cluster not just single node
* topic with multiple paritions then use consumer group on it
* HTTP/2
* retries on LeaderNotAvailable errors
* Kafka consistency patterns

## Curls

```bash
curl localhost:19092/topics
curl 'http://localhost:19092/topics/topic2/messages?group=1'
```

## Kafkacat

run from jump

```bash
kafkacat -b 10.96.1.1:9092 -L
kafkacat -b 10.96.1.1:9092 -t topic1 -P
```
