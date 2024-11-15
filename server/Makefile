include .env

MIGRATIONS_PATH = ./internal/database/migrations
DB_URL := "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)"

.PHONY: create
create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(name)

.PHONY: up
up:
	@migrate -database=$(DB_URL) -path=$(MIGRATIONS_PATH) up

.PHONY: down
down:
	@migrate -database $(DB_URL) -path $(MIGRATIONS_PATH) down

