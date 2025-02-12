build:
	@go build -o bin/ecom-example cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom-example
