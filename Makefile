.PHONY: build install clean

APP_NAME=gituser
BIN_PATH=$(HOME)/bin

build:
	go build -o $(APP_NAME)

install: build
	mkdir -p $(BIN_PATH) && mv $(APP_NAME) $(BIN_PATH)
	
clean: $(APP_NAME)
	rm $(APP_NAME)
