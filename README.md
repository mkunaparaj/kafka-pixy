# kafka-pixy-POC

## Usage

expects helm + kubernetes running locally

```bash
make run
make ui # creates cluster pointed to zookeepers deployed in `make run`
make clean # to destroy everything
```

## TODO

* form kafka cluster not just single node
* topic with multiple paritions then use consumer group on it
* retries on LeaderNotAvailable errors
* multiple kafka clusters
* seem to be losing messages in between start up and down
* leverage context for goroutine management
* partition scheme

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

## Observations

multiple consumers on same partition
proxy dies, maintains consumer group offset state on return
set limits on unacknowledged requests
the proxy will batch messages and then async pass to kafka to increase speed
need a way to gracefully stop consumers to not leave a message half processed