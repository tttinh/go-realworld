# Print warnings if the .env does not exist.
ifeq (,$(wildcard .env))
$(info WARN: .env file not found!)
$(info WARN: some environment variables may need to be set!)
endif

-include .env


###############################################################################
# CODE GENERATION
###############################################################################

# Generate code from .sql files for data layers.
sqlc:
	sqlc generate
.PHONY: sqlc

###############################################################################
# MIGRATION
###############################################################################

# Create new migration files.
# Example: make add-migration name=create_table_abc
add-migration:
	migrate create -ext sql -seq -dir $(POSTGRES_MIGRATION) $(name)
.PHONY: add-migration


# Run database migration up.
up-migration:
	migrate -path $(POSTGRES_MIGRATION) -database $(POSTGRES_URI) up
.PHONY: up-migration


# Run database migration down.
down-migration:
	migrate -path $(POSTGRES_MIGRATION) -database $(POSTGRES_URI) down -all
.PHONY: down-migration


###############################################################################
# TEST
###############################################################################

# Run E2E test.
test-e2e:
	./scripts/e2e.sh
.PHONY: test-e2e


###############################################################################
# DOCKER
###############################################################################

compose-up:
	docker compose up
.PHONY: compose-up

compose-down:
	docker compose down
.PHONY: compose-down

# Rebuild backend image for docker compose.
compose-build:
	docker build . -t conduit-backend
.PHONY: compose-build


###############################################################################
# OTHERS
###############################################################################

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
