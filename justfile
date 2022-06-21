#!/usr/bin/env just --justfile

run:
    GOOS=windows GOARCH=amd64 CGO_ENABLED=1  go run .
    #GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -x ./

update:
  go get -u
  go mod tidy -v