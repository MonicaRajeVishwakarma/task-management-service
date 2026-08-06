.PHONY: help test test-cover coverage fmt vet tidy clean

help:
	@echo "Available commands"
	@echo "  make test       - Run all tests"
	@echo "  make test-cover - Run tests with coverage percentage"	
	@echo "  make coverage   - Generate test coverage report"
	@echo "  make fmt        - Format Go code"
	@echo "  make vet        - Run Go static analysis"
	@echo "  make tidy       - Clean go dependencies"
	@echo "  make clean      - Remove generated files"

test:
	go test ./...

test-cover:
	go test ./... -cover

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -f coverage.out


