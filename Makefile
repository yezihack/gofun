#!/usr/bin/env bash

all: dev run

dev:clean fmt build
	./run/gofun -c gofun.toml

run:clean build
	./run/gofun

deam:
	nohup ./run/gofun > /dev/null 2>&1 &

kill:
	pid=$(shell ps -ef |grep -v grep |grep gofun |awk {'print $2'})
	echo $pid
fmt:
	gofmt -l -w ./

build:
	go build -v -o ./run/gofun ./app

clean:
	rm -rf run/gofun

vendor:
	govendor remove +u
	govendor add +e

copy:
	cp run/config.toml.simple run/config.toml
