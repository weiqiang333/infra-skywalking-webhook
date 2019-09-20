#!/usr/bin/env bash

export GOARCH=amd64
export GOOS=linux
export GCCGO=gc

go build -o bin/infra-skywalking-webhook main.go