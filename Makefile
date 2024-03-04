test:
	go test -v -race ./...

install-gofumpt:
	go install mvdan.cc/gofumpt@latest

install-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2

format: ## run go formatter
	gofumpt -l -w .
lint:
	@which golangci-lint || make install-golangci-lint
	golangci-lint run --out-format=github-actions --timeout=10m