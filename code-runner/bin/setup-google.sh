#!/usr/bin/env bash
export DOCKER_CERT_PATH="$(pwd)/remote-daemon/client"
export DOCKER_HOST="https://127.0.0.1:6666"
go run main.go --config env/gcp-config.json



