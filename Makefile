## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'


## init: install requirements
.PHONY: init
init:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
	mv migrate.linux-amd64 $GOPATH/bin/migrate
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: dependencies
dependencies:
	go mod download

# ==================================================================================== #
# APPLICATIONS
# ==================================================================================== #

## runweb: run the web application
.PHONY: runweb
runweb:
	go run ./cmd/web -db-dsn=${DB_DSN}

## runapi: run the web application
.PHONY: runapi
runapi:
	go run ./cmd/api -db-dsn=${DB_DSN}

# ==================================================================================== #
# MIGRATIONS
# ==================================================================================== #

.PHONY: migrate
migrate:
	migrate -path=./migrations -database="${DB_DSN}" up

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
