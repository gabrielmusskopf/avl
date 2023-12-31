dev:
	go run ./cmd/main.go ./cmd/cmd.go

build: 
	mkdir -p ./bin
	go build -o ./bin/main ./cmd/*.go 

run: 
	./bin/main

tests:
	@go test -v ./pkg
	@go test -v ./pkg/types

todo:
	@grep -rn --exclude "Makefile" TODO | grep -oP '//\K.*'

http: 
	./bin/main -http

start: build
	./bin/main
