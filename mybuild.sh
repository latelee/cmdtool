#!/bin/sh

GO111MODULE=on go build -mod vendor -o cmdtool.exe main.go