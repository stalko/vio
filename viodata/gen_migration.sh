#!/bin/sh

MIGRATION_NAME=$1

if [ -z "$MIGRATION_NAME" ]
then
    echo "usage: ./gen_migration.sh [migration_name]"
    exit 1
fi

MIGRATIONS_DIR="$(pwd)/migrations"

docker run --rm -v "$MIGRATIONS_DIR":/migrations migrate/migrate create -ext sql -dir /migrations/ $MIGRATION_NAME