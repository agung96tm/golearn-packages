## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: init
## init: additional requirements
init:
	go install github.com/hibiken/asynq/tools/asynq@latest

dependencies:
## dependencies: install dependencies
	go mod download

# ==================================================================================== #
# APPLICATIONS
# ==================================================================================== #

## runweb: run the web application
.PHONY: runweb
runweb:
	go run ./cmd/web

## runapi: run the web application
.PHONY: runapi
runapi:
	go run ./cmd/api

# ==================================================================================== #
# WORKERS & SCHEDULERS
# ==================================================================================== #

## workerweb: run worker for web application
.PHONY: workerweb
workerweb:
	go run ./cmd/web -serve=worker

## workerapi: run worker for api application
.PHONY: workerapi
workerapi:
	go run ./cmd/api -serve=worker

## workerweb: run scheduler for web application
.PHONY: schedulerweb
schedulerweb:
	go run ./cmd/web -serve=scheduler

## schedulerapi: run scheduler for api application
.PHONY: schedulerapi
schedulerapi:
	go run ./cmd/api -serve=scheduler

# ==================================================================================== #
# BUILD
# ==================================================================================== #

current_time = $(shell date --iso-8601=seconds)
linker_flags = '-s -X main.buildTime=${current_time}'

## buildweb: build the web application
.PHONY: buildweb
buildweb:
	@echo "Building web..."
	go build -ldflags=${linker_flags} -o=./bin/web ./cmd/web
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/web ./cmd/web

## buildapi: build the web application
.PHONY: buildapi
buildapi:
	@echo "Building api..."
	go build -ldflags=${linker_flags} -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/api ./cmd/api
