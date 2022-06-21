#!/usr/bin/env just --justfile
env := "w"

prefix := if env == "w" { "GOOS=windows GOARCH=amd64 CGO_ENABLED=1" } else { "GOOS=linux GOARCH=amd64" }

run:
    {{prefix}}  go run .

build:
    {{prefix}}  go build .

test:
    {{prefix}} go test ./test

update:
  go get -u
  go mod tidy -v