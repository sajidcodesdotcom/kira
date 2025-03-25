include /$(PWD)/.env
build:
	go build -o ./bin/server ./cmd/server/server.go

run: build
	./bin/server

test: 
	go test ./...

lint: 
	golangci-lint run --enable-all

migrate-up:
	goose -dir migrations postgres ${DATABASE_URL} up

migrate-down:
	goose -dir migrations postgres ${DATABASE_URL} down

migrate-status:
	goose -dir migrations postgres ${DATABASE_URL} status

migrate-version:
	goose -dir migrations postgres ${DATABASE_URL} version


migrate-create:
	@read -p "Enter migration name: " name; \
	goose -dir migrations create $$name sql