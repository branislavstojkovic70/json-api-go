build:
	@go build -o bin/json-api-go

run: build
	@./bin/json-api-go

test:
	@go test -v ./...