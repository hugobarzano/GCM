#!/usr/bin/env bash
export DOCKER_CERT_PATH="$(pwd)/remote-daemon/client"
export DOCKER_HOST="https://35.233.23.148:5555"
go run main.go --config env/dev-config.json



