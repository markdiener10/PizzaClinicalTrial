SHELL := /usr/bin/env bash

.DEFAULT_GOAL := test

.PHONY: clean test coverage

clean:
	go clean -testcache -cache -modcache
	@echo "Project CLEANED ##############"

test: 
	go vet 
	go test 

coverage: 	
	rm -f ./coverage.out
	go test -coverprofile=./coverage.out
	go tool cover -html=./coverage.out
