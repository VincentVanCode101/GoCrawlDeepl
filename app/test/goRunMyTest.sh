#!/usr/bin/env bash

mkdir -p coverage

go test .././... -coverprofile=coverage/test.out

go tool cover -html=./coverage/test.out -o coverage/coverage.html