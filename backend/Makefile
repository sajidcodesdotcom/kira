build:
	go build -o ./bin/server ./cmd/server/server.go

run: build
	./bin/server

test: 
	go test ./...

lint: 
	golangci-lint run --enable-all