# Check and load the .env file
ifeq (,$(wildcard .env))
$(error .env file not found)
endif

include .env
export $(shell sed 's/=.*//' .env)


# Start a Postgres database in Docker.
start-db:
	docker run --rm -d --name conduit-db -p 5432:5432 \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
 		postgres:16-alpine
.PHONY: start-db


# Stop the Postgres database in Docker.
stop-db:
	docker stop conduit-db
.PHONY: stop-db


# Create new migration files.
# Usage: make name=update_column migrate-file
migrate-file:
	migrate create -ext sql -seq -dir $(POSTGRES_MIGRATION) $(name)
.PHONY: migrate-file


# Run database migration up.
migrate-up:
	migrate -path $(POSTGRES_MIGRATION) -database $(POSTGRES_URI) up
.PHONY: migrate-up


# Run database migration down.
migrate-down:
	migrate -path $(POSTGRES_MIGRATION) -database $(POSTGRES_URI) down -all
.PHONY: migrate-down