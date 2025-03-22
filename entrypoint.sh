#!/bin/sh
set -e

echo "Waiting for MinIO to be ready..."
until curl -s http://ecommerce.minio:9000/minio/health/live; do
  sleep 2
done

echo "Retrieving MinIO credentials..."
export MINIO_ACCESSKEY="minioadmin"
export MINIO_SECRETKEY="minioadmin123"

echo "Starting application..."
exec "/app"

