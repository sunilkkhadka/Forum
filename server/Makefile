include .env

# Variables
MIGRATIONS_PATH = ./internal/database/migrations
DB_URL := "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)"

DOCKER_IMAGE_NAME="forum_image"
DOCKER_DATABASE_HOST="$(DB_HOST)"

# Migration Commands
.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(name)

.PHONY: migrate-up
migrate-up:
	@migrate -database=$(DB_URL) -path=$(MIGRATIONS_PATH) up

.PHONY: migrate-down
migrate-down:
	@migrate -database $(DB_URL) -path $(MIGRATIONS_PATH) down
	

# Docker Commands
.PHONY: build-nocache
build-nocache:
	@sudo docker build --no-cache -t $(DOCKER_IMAGE_NAME) .

.PHONY: build
build:
	@sudo docker build -t $(DOCKER_IMAGE_NAME) .

.PHONY: list
list: 
	@sudo docker ps

.PHONY: list-all
list-all:
	@sudo docker ps -a

.PHONY: up
up:
	@sudo docker compose up $(name) -d

.PHONY: down
down:
	@sudo docker compose down $(name)

.PHONY: up-all
up-all:
	@sudo docker compose up -d

.PHONY: down-all
down-all:
	@sudo docker compose down

.PHONY: remove-container
remove-container:
	@sudo docker rm $(name)

.PHONY: restart
restart:
	@sudo docker compose restart

.PHONY: remove-image
remove-image:
	@sudo docker rmi $(name)

.PHONY: logs
logs:
	@sudo docker logs -f $(name)

.PHONY: exec
exec:
	@sudo docker exec -it $(name) bash

