# Print warnings if the .env does not exist.
ifeq (,$(wildcard .env))
$(info WARN: .env file not found!)
$(info WARN: some environment variables may need to be set!)
endif

-include .env

# Generate code from .sql files into infra folder.
sqlc:
	sqlc generate
.PHONY: sqlc


# Create new migration files.
# Example: make add-migration name=create_table_abc
add-migration:
	migrate create -ext sql -seq -dir $(MK_DB_MIGRATION) $(name)
.PHONY: add-migration


# Run database migration up.
up-migration:
	migrate -path $(MK_DB_MIGRATION) -database $(MK_DB_URI) up
.PHONY: up-migration


# Run database migration down.
down-migration:
	migrate -path $(MK_DB_MIGRATION) -database $(MK_DB_URI) down -all
.PHONY: down-migration

# Run E2E test.
e2e-test:
	./scripts/e2e.sh
.PHONY: e2e-test