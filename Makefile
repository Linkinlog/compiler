lint:
	@go mod tidy
	@gofumpt -d -w .
	@golangci-lint run

test:
	@go test -v ./...

.PHONY: lint test
