#!/bin/sh
until PGPASSWORD=$POSTGRES_PASSWORD psql -h "$POSTGRES_HOST" -U "$POSTGRES_USER" -d postgres -c '\q'; do
    >&2 echo "PostgreSQL is still unavailable - sleeping"
    sleep 1
done

>&2 echo "PostgreSQL is up - continuing"

RESULT=$(PGPASSWORD=$POSTGRES_PASSWORD psql -h "$POSTGRES_HOST" -U "$POSTGRES_USER" -d "$POSTGRES_DB" -t -c "SELECT 1 FROM pg_database WHERE datname = 'segmentation'")
if [ -z "$RESULT" ]; then
    echo "Database does not exist, creating..."
    PGPASSWORD=$POSTGRES_PASSWORD psql -h "$POSTGRES_HOST" -U "$POSTGRES_USER" -c "CREATE DATABASE segmentation"
    echo "Database migration complete!"
else
    echo "Database exists, skipping creation..."
    ./dentistMigrate
    ./patientMigrate
    ./segmentationMigrate
    echo "Database migration complete!"
fi
exec "$@"