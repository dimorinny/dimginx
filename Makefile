GOPATH := $(shell pwd)

all:
	rm -f ./httpd
	go test ./...
	go build -o ./httpd ./src/main.go
	