#!make

build:
	@go build -v .

test:
	@go test -v ./... -race

coverage:
	@go test ./... -cover -race -coverprofile=coverage.out -covermode=atomic
	@go tool cover -html=coverage.out -o coverage.html

fmt:
	@gofmt -w .

clean:
	@rm -rf coverage.html coverage.out
