dev:
	go run ./cmd/main.go ./cmd/cmd.go

build: 
	mkdir -p ./bin
	go build -o ./bin/main ./cmd/*.go 

run: 
	./bin/main

run-http: 
	./bin/main -http

start: build
	./bin/main

start-http: 
	mkdir -p ./bin
	go build -o ./bin/main ./cmd/*.go 
	./bin/main -http
