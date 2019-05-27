#!/usr/bin/env bash

dev:clean fmt build

run:clean build
	./run/gofun

fmt:
	gofmt -l -w ./

build:
	go build -v -o ./run/gofun ./app

clean:
	rm -rf run/*

vendor:
	govendor add +e
	govendor remove +u
