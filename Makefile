.PHONY: build install clean fmt lint

APP_NAME=gituser
BIN_PATH=$(HOME)/bin

fmt:
	gofmt -s -l -w .

lint: fmt
	golangci-lint run

build:
	go build -o $(APP_NAME)

install: build
	mkdir -p $(BIN_PATH) && mv $(APP_NAME) $(BIN_PATH)

clean: $(APP_NAME)
	rm $(APP_NAME)

test-release:
	 goreleaser release --snapshot --skip-publish --rm-dist
