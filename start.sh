#!/bin/sh

set -e

echo "run db migration"
source /app/app.env
./goose up -dir /app/migration 

echo "start the app"
exec "$@"