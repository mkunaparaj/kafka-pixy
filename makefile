SHELL := /bin/bash
HELM_DEPLOY_NAME=events

.PHONY: clean deploy

run:
	helm install $(HELM_DEPLOY_NAME) deploy

ui:
	curl -s -X POST localhost:9000/clusters \
	--data "name=main_cluster&zkHosts=10.96.1.2:2181&kafkaVersion=2.2.0&jmxEnabled=true&pollConsumers=true&activeOffsetCacheEnabled=true&tuning.brokerViewUpdatePeriodSeconds=30&tuning.clusterManagerThreadPoolSize=2&tuning.clusterManagerThreadPoolQueueSize=100&tuning.kafkaCommandThreadPoolSize=2&tuning.kafkaCommandThreadPoolQueueSize=100&tuning.logkafkaCommandThreadPoolSize=2&tuning.logkafkaCommandThreadPoolQueueSize=100&tuning.logkafkaUpdatePeriodSeconds=30&tuning.partitionOffsetCacheTimeoutSecs=5&tuning.brokerViewThreadPoolSize=4&tuning.brokerViewThreadPoolQueueSize=1000&tuning.offsetCacheThreadPoolSize=4&tuning.offsetCacheThreadPoolQueueSize=1000&tuning.kafkaAdminClientThreadPoolSize=4&tuning.kafkaAdminClientThreadPoolQueueSize=1000&tuning.kafkaManagedOffsetMetadataCheckMillis=30000&tuning.kafkaManagedOffsetGroupCacheSize=1000000&tuning.kafkaManagedOffsetGroupExpireDays=7&securityProtocol=PLAINTEXT&saslMechanism=DEFAULT"

clean:
	helm uninstall $(HELM_DEPLOY_NAME)