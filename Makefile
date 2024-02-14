## runweb: run the web application
.PHONY: runweb
runweb:
	go run ./cmd/web


## buildweb: build the web application
.PHONY: buildweb
buildweb:
	@echo "Building web..."
	go build -ldflags=${linker_flags} -o=./bin/web ./cmd/web
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/web ./cmd/web


## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
