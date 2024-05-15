.PHONY: build
build:
	${MAKE} -C apps build
	@mkdir -p bin
	@cp apps/ai-bots-be/bin/aibots-service bin/aibots-service
	@echo ""
	@echo "Build is done 🎉"
	@echo "Run aibots-service with ./bin/aibots-service"
	@echo ""


.PHONY: test
test:
	${MAKE} -C apps test