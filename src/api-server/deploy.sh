#!/bin/bash

echo '[deploy] Build dist ...'
env GOOS=linux GOARCH=amd64 go build -v -o ./dist/server gitlab.com/hs-api-go
echo '[deploy] Build dist DONE'
