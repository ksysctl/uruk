include configs/.env

ENV_FILE=configs/.env
CONFIG_FILE=deployments/docker-compose.yml
PROJECT_NAME=$(notdir $(PWD))
BASE_ARGS=-f $(CONFIG_FILE) --project-name $(PROJECT_NAME) --env-file $(ENV_FILE)

.PHONY: \
	build up down logs clean deps

all: build
build:
	docker-compose $(BASE_ARGS) up --build
rebuild:
	docker-compose $(BASE_ARGS) build --no-cache
up:
	docker-compose $(BASE_ARGS) up -d
down:
	docker-compose $(BASE_ARGS) down
logs:
	docker-compose $(BASE_ARGS) --verbose logs -f
clean:
	gofmt -w .
	go clean
deps:
	go mod tidy -v