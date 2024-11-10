.PHONY: build
build:
	${MAKE} -C ai-chats-fe build
	@cp -r ai-chats-fe/dist ai-chats-be/web
	${MAKE} -C ai-chats-be build

.PHONY: tests
tests:
	${MAKE} -C ai-chats-be tests
	${MAKE} -C ai-chats-fe tests

.PHONY: test-integration
test-integration:
	${MAKE} -C ai-chats-be test-integration

.PHONY: tests-e2e
tests-e2e:
	${MAKE} -C ai-chats-e2e tests-e2e

.PHONY: format
format:
	${MAKE} -C ai-chats-be format
	${MAKE} -C ai-chats-fe format