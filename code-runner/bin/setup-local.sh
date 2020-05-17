#!/usr/bin/env bash
export DOCKER_CERT_PATH="$(pwd)/env/local/remote-daemon/client"
export DOCKER_HOST="https://dev.gcm-coderunner.com:5555"
go run main.go --config env/local/local-config.json



