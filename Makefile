all:	
	export GOPATH=${PWD}	
	rm -f ./httpd
	go build -o ./httpd ./src/main.go	