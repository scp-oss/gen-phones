#!/usr/bin/env sh
# Builds the gen-phones Docker image, taking PORT from .env (falls back to
# 8080 if .env is missing or PORT isn't set there). Usage: ./build.sh [tag]
set -eu

cd "$(dirname "$0")"

if [ -f .env ]; then
  # shellcheck disable=SC1091
  set -a
  . ./.env
  set +a
fi

PORT="${PORT:-8080}"
TAG="${1:-gen-phones:latest}"

echo "Building $TAG with PORT=$PORT (from .env)"
docker build --build-arg PORT="$PORT" -t "$TAG" .
