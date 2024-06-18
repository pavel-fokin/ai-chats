.PHONY: build
build:
	@echo "Building..."
	@go build -o bin/aibots-service cmd/aibots-service/main.go

.PHONY: run
run:
	@go run cmd/aibots-service/main.go

.PHONY: test
test:
	@echo "Running unit tests..."
	@go test -count 1 -cover -timeout 5s ./...
	@echo "✅ All unit tests are good"

.PHONY: test-integration
test-integration:
	@echo "Running integration tests..."
	@go test -count 1 -cover -tags=integration ./...
	@echo "✅ All integration tests are good"

.PHONY: test-race
test-race:
	@echo "Running tests to check data racing..."
	@go test -count 1 -cover -race -tags=race --timeout 5s ./...
	@echo "✅ All race tests are good"

.PHONY: test-all
test-all: test test-integration test
	@echo "✅ All tests are good"
	@echo ""

.PHONY: format
format:
	@gofmt -w internal/ cmd/
