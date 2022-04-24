mkSOURCE_PATH := $(PWD)/cmd/main.go
BIN_PATH := $(PWD)/bin/main
EVENT_FILE := event.json

# Build source code and create executable or image
build:
	@echo "building source code..."
	env GOOS=linux GOARCH=amd64 go build -o $(BIN_PATH) $(SOURCE_PATH)

# Run locally
run: build
	@echo "running..."
	sam local invoke Function -e $(EVENT_FILE)

start: build
	@echo "starting..."
	sam local start-api
