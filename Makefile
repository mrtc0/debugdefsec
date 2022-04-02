.PHONY: build
build:
	mkdir -p build
	go build -tags netgo -ldflags '-w -s -extldflags "-static"' -o build/debugdefsec main.go
