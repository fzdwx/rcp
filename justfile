#!/usr/bin/env just --justfile

prefix := "GOOS=windows GOARCH=amd64 CGO_ENABLED=1"

build:
    {{prefix}}  go build .

run:
    {{prefix}}  go run .
    #GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -x ./

test:
    {{prefix}} go test ./test

update:
  go get -u
  go mod tidy -v