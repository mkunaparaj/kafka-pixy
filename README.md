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
* multiple kafka clusters
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
can reset consumer group offsets, set init offset for a new consumer group. Can only do this to non existing consumer groups,
existing or "in use" consumer groups wont be affected by this command.
noisy error logs when a message isn't ack'd immediately in pixy logs
