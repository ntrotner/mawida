#!make
TARGET ?= prod
include ${PWD}/environment/global.env

# Env setup
DOCKER_COMPOSE := ""
ifneq ($(shell command -v docker-compose),)
	DOCKER_COMPOSE = docker-compose
else ifneq ($(shell command -v docker compose),)
	DOCKER_COMPOSE = docker compose
else
	$(error Neither docker-compose nor docker compose is available)
	exit 1
endif

SHELL := /bin/bash

setup-env:
	cp ${PWD}/scripts/git/* ${PWD}/.git/hooks/
	go install github.com/gordonklaus/ineffassign@latest
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	go install github.com/bykof/go-plantuml@latest
	cd ./src/ui && npm install

# TARGET = dev/prod
setup-env-file:
	./scripts/setup-env-file.sh ${TARGET}
  export $(shell sed 's/=.*//' ${PWD}/environment/global.env)

setup-env-file-deployment:
	./scripts/setup-env-file.sh prod
  export $(shell sed 's/=.*//' ${PWD}/environment/global.env)

deploy:
	make clean || true
	docker network create host_network || true
	export docker_env=PROD && ${DOCKER_COMPOSE} --profile deployment up --build --remove-orphans --force-recreate -d

# Start
# run app in local for development
# TARGET = dev/prod
run:
	make clean || true
	docker network create host_network || true
	make setup-env-file
	export docker_env=${TARGET} && ${DOCKER_COMPOSE} --profile local up --build --remove-orphans --force-recreate

clean:
	make setup-env-file-deployment
	export docker_env=${TARGET} && $(DOCKER_COMPOSE) --profile local down

clean-deployment:
	make setup-env-file-deployment
	export docker_env=prod && $(DOCKER_COMPOSE) --profile deployment down

# OpenAPI
openapi-generate-backend:
	docker run --rm --user 1000:1000 \
			-v ${PWD}/src/backend/open-api/temp:/local \
			-v ${PWD}/shared:/shared \
			-v /usr/local/go/bin/gofmt:/usr/local/bin/gofmt \
			-e GO_POST_PROCESS_FILE="/usr/local/bin/gofmt -w" \
			openapitools/openapi-generator-cli generate \
			-i /shared/openApi.yaml \
			-g go-server \
			-o /local/ \
			--additional-properties="outputAsLibrary=true,enablePostProcessFile=true"
	cat ${PWD}/shared/openApi.yaml | yq -o=json -I2 > ${PWD}/src/backend/open-api/swaggerui/swagger.json

openapi-generate-ui:
	docker run --rm --user 1000:1000 \
			-v ${PWD}/src/ui/src/lib/open-api:/local \
			-v ${PWD}/shared:/shared \
			openapitools/openapi-generator-cli generate \
			-i /shared/openApi.yaml \
			-g typescript-fetch \
			-o /local/

openapi-generate:
	make openapi-generate-backend
	make openapi-generate-ui
