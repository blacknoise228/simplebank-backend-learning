#!/bin/sh

set -e

echo "run db migration"
./goose up -dir /app/migration 

echo "start the app"
exec "$@"