.PHONY: build
build:
	@echo "Building..."
	@go build -o bin/aibots-service cmd/aibots-service/main.go

.PHONY: test
test:
	@echo "Running unit tests..."
	@go test -cover ./...