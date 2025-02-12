#!/bin/bash

MIGRATION_PATH="../infra/postgres/migration"
DB_URI="postgres://tinhtt:tinhtt@localhost:5432/conduit?sslmode=disable"

if [ "$1" = "up" ]; then
    migrate -path $MIGRATION_PATH -database $DB_URI up
    exit 0
fi

if [ "$1" = "down" ]; then
    migrate -path $MIGRATION_PATH -database $DB_URI down
    exit 0
fi

echo "Usage: ./migrate.sh <up | down>"