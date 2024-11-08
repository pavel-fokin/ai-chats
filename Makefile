.PHONY: build
build:
	${MAKE} -C apps build
	@mkdir -p bin
	@cp apps/ai-chats-be/bin/ai-chats-service bin/ai-chats-service
	@echo ""
	@echo "Build is done 🎉"
	@echo "Run ai-chats-service with ./bin/ai-chats-service"
	@echo ""


.PHONY: tests
tests:
	${MAKE} -C apps tests

.PHONY: test-integration
test-integration:
	${MAKE} -C apps test-integration

.PHONY: tests-e2e
tests-e2e:
	${MAKE} -C apps tests-e2e

.PHONY: test-all
test-all: test test-integration test-e2e
	@echo "✅ All tests are good"
	@echo ""

.PHONY: format
format:
	${MAKE} -C apps format