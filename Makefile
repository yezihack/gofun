#!/usr/bin/env bash

all: dev run

dev:clean fmt build


run:clean build
	./run/gofun

fmt:
	gofmt -l -w ./

build:
	go build -v -o ./run/gofun ./app

clean:
	rm -rf run/gofun

vendor:
	govendor add +e
	govendor remove +u
copy:
	cp run/config.toml.simple run/config.toml
