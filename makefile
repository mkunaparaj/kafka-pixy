SHELL := /bin/bash
HELM_DEPLOY_NAME=events

.PHONY: clean deploy

run:
	helm install $(HELM_DEPLOY_NAME) deploy

clean:
	helm uninstall $(HELM_DEPLOY_NAME)