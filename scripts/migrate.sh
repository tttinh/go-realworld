#!/bin/bash

MIGRATION_PATH="file://db/postgres/migration"
DB_URI="postgres://tinhtt:tinhtt@localhost:5432/conduit?sslmode=disable"

if [ "$1" = "up" ]; then
    migrate -source $MIGRATION_PATH -database $DB_URI up
    exit 0
fi

if [ "$1" = "down" ]; then
    migrate -source $MIGRATION_PATH -database $DB_URI down -all
    exit 0
fi

echo "Usage: ./migrate.sh <up | down>"


migrate -source file://db/postgres/migration -database "postgres://tinhtt:tinhtt@localhost:5432/conduit?sslmode=disable" up