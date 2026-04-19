NPM=npm
GO=go

.PHONY: install
install:
	cd src/frontend && $(NPM) install
	cd src/backend && $(GO) mod tidy

.PHONY: test
test:
	cd src/backend && $(GO) test ./...

.PHONY: build
build: install test
	cd src/frontend && $(NPM) run build
	mv src/frontend/dist src/backend/embed/webui
	cd src/backend && $(GO) build
	mkdir dist
	mv src/backend/dlp-ui dist

.PHONY: clean
clean:
	rm -rf src/frontend/node_modules
	rm -rf src/backend/embed/webui
	rm -rf dist