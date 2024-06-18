.PHONY: build
build:
	${MAKE} -C aibots-fe build
	@cp -r aibots-fe/dist aibots-be/web
	${MAKE} -C aibots-be build

.PHONY: test
test:
	${MAKE} -C aibots-be test
	${MAKE} -C aibots-fe test

.PHONY: test-integration
test-integration:
	${MAKE} -C aibots-be test-integration

.PHONY: test-e2e
test-e2e:
	# ${MAKE} -C test test-e2e
	${MAKE} -C playwright test-e2e

.PHONY: format
format:
	${MAKE} -C aibots-be format
	${MAKE} -C aibots-fe format