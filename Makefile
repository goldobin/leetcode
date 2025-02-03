default: fmt lint test

fmt:
	@goimports -local "github.com/goldobin/leetcode" -l -w .
	@gofumpt -l -w .

lint:
	@golangci-lint run ./...

test:
	@go test ./...