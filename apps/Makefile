.PHONY: build
build:
	${MAKE} -C aibots-fe build
	@cp -r aibots-fe/dist aibots-be/internal/web
	${MAKE} -C aibots-be build

.PHONY: test
test:
	${MAKE} -C aibots-be test
	${MAKE} -C aibots-fe test