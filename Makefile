GO_BUILD_TARGET=./src

.PHONY: source
source:
	go mod tidy
	go build -v $(GO_BUILD_TARGET)
