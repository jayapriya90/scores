GO111MODULES=on

.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	# go tool dist list command shows a list of supported OS/Arch combinations
	GOOS=darwin GOARCH=amd64 go build -o scores github.com/jayapriya90/scores

.PHONY: run
## run: build and run the application
run:
	go run scores.go

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning binaries..."
	go clean

.PHONY: fmt
## fmt: format Go source code
fmt:
	go fmt ./...

.PHONY: lint
## lint: lint Go source code
lint:
	golint ./...

.PHONY: test
## test: run Go tests
test:
	go test ./...

.PHONY: tidy
## tidy: cleanup unused dependencies
tidy:
	go mod tidy

.PHONY: vet
## vet: check for common errors in Go source code
vet:
	go vet ./...

.PHONY: help
## help: prints the help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
