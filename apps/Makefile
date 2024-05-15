.PHONY: build
build:
	${MAKE} -C ai-bots-fe build
	@cp -r ai-bots-fe/dist ai-bots-be/internal/web
	${MAKE} -C ai-bots-be build

.PHONY: test
test:
	${MAKE} -C ai-bots-be test
	${MAKE} -C ai-bots-fe test