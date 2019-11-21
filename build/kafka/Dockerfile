FROM debian:latest

# install dependencies
RUN apt update
RUN apt install -y wget default-jdk procps tree

# install kafka version 2.3.0
RUN wget http://mirrors.ocf.berkeley.edu/apache/kafka/2.3.0/kafka_2.12-2.3.0.tgz
RUN mkdir kafka && tar -xzvf kafka_2.12-2.3.0.tgz -C kafka --strip-components 1

# init scripts
COPY start.sh /

ENTRYPOINT ["./start.sh"]
