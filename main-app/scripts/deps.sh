#!/usr/bin/env bash

go mod tidy -v && go mod download && go mod vendor;