GITHUB_REF_NAME ?= dev
GITHUB_SHA ?= none

NPM := npm

GO := go
CGO_ENABLED := 0

.PHONY: install-frontend
install-frontend:
	cd src/frontend && $(NPM) install

.PHONY: install-backend
install-backend:
	cd src/backend && $(GO) mod download

.PHONY: install
install: install-frontend install-backend

.PHONY: test-frontend
test-frontend:
	cd src/frontend && $(NPM) run lint

.PHONY: test-backend
test-backend:
	cd src/backend && $(GO) test ./...

.PHONY: test
test: test-frontend test-backend

.PHONY: build-frontend
build-frontend:
	cd src/frontend && $(NPM) run build

.PHONY: build-backend
build-backend:
	mv src/frontend/dist src/backend/embed/webui

	cd src/backend && $(GO) build -trimpath -ldflags="-s -w" -ldflags "-X dlp-ui/cmd.tag=${GITHUB_REF_NAME} -X dlp-ui/cmd.commit=${GITHUB_SHA}" -o dlp-ui

	mkdir -p dist
	mv src/backend/dlp-ui dist

.PHONY: build
build: build-frontend build-backend

.PHONY: clean
clean:
	rm -rf src/frontend/node_modules
	rm -rf src/frontend/dist
	rm -rf src/backend/embed/webui
	rm -rf dist