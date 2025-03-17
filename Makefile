# Check and load the .env file
ifeq (,$(wildcard .env))
$(error .env file not found)
endif

include .env
export $(shell sed 's/=.*//' .env)

# Generate code from .sql files into infra folder.
sqlc:
	sqlc generate
.PHONY: sqlc


# Create new migration files.
# Example: make migrate-file name=create_table_abc
migrate-file:
	migrate create -ext sql -seq -dir $(MK_DB_MIGRATION) $(name)
.PHONY: migrate-file


# Run database migration up.
migrate-up:
	migrate -path $(MK_DB_MIGRATION) -database $(MK_DB_URI) up
.PHONY: migrate-up


# Run database migration down.
migrate-down:
	migrate -path $(MK_DB_MIGRATION) -database $(MK_DB_URI) down -all
.PHONY: migrate-down

# Run E2E test.
e2e:
	./scripts/e2e.sh
.PHONY: e2e