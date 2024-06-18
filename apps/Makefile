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
	${MAKE} -C aibots-be build
	@echo "✅ run ai service in background and get the pid"
	@nohup aibots-be/bin/aibots-service > /dev/null 2>&1 & echo $$! > aibots-be.pid
	@echo "✅ wait for service to start"
	@sleep 1
	${MAKE} -C test test-e2e
	@echo "✅ kill the service "
	@kill -9 `cat aibots-be.pid`
	@rm aibots-be.pid
	@echo "✅ All e2e tests are good"

.PHONY: format
format:
	${MAKE} -C aibots-be format
	${MAKE} -C aibots-fe format