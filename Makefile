default: test

fmt:
	@gofumpt -l -w .

test:
	@go test ./...