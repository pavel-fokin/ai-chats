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

.PHONY: test-e2e
test-e2e:
	# ${MAKE} -C test test-e2e
	${MAKE} -C playwright test-e2e

.PHONY: format
format:
	${MAKE} -C ai-chats-be format
	${MAKE} -C ai-chats-fe format