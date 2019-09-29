lint:
	golangci-lint run

format:
		goimports -w -l .
		go fmt

test:
		go test --bench .

all: format lint test