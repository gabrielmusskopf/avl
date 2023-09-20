dev:
	go run ./cmd/main.go

build: 
	mkdir -p ./bin
	go build -o ./bin/main ./cmd/main.go 

run: 
	./bin/main

start: build
	./bin/main
