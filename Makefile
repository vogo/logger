lint:
	golangci-lint run

format:
		goimports -w -l .
		go fmt

test:
		go test --bench .

all: format lint test

bench_logger:
	go test logger_test.go logger.go --bench .

bench_logger_prof:
	go test logger_test.go logger.go --bench . -cpuprofile cpu.prof  -memprofile mem.prof