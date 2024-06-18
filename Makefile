.PHONY: build
build:
	${MAKE} -C apps build
	@mkdir -p bin
	@cp apps/aibots-be/bin/aibots-service bin/aibots-service
	@echo ""
	@echo "Build is done 🎉"
	@echo "Run aibots-service with ./bin/aibots-service"
	@echo ""


.PHONY: test
test:
	${MAKE} -C apps test

.PHONY: test-integration
test-integration:
	${MAKE} -C apps test-integration

.PHONY: test-e2e
test-e2e:
	${MAKE} -C apps test-e2e

.PHONY: test-all
test-all: test test-integration test-e2e
	@echo "✅ All tests are good"
	@echo ""

.PHONY: format
format:
	${MAKE} -C apps format