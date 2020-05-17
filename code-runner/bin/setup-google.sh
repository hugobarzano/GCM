#!/usr/bin/env bash
export DOCKER_CERT_PATH="$(pwd)/env/google/remote-daemon/client"
export DOCKER_HOST="https://127.0.0.1:5555"
#go run main.go --config env/google/gcp-config.json
./code-runner.linux  --config env/google/gcp-config.json


