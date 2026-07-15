NPM := npm

GO := go
CGO_ENABLED := 0

UPX := upx

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
	mkdir -p src/backend/embed/frontend

	cd src/backend && $(GO) test ./...

.PHONY: test
test: test-frontend test-backend

.PHONY: build-frontend
build-frontend:
	cd src/frontend && $(NPM) run build

	mkdir -p dist
	mv src/frontend/dist dist
	mv dist/dist dist/frontend

.PHONY: build-backend
build-backend:
	mkdir -p src/backend/embed
	cp -r dist/frontend src/backend/embed

	cd src/backend && $(GO) build -trimpath -ldflags="-s -w" -o dlp-ui
	-cd src/backend && $(UPX) --best --lzma dlp-ui

	mkdir -p dist/backend
	mv src/backend/dlp-ui dist/backend

.PHONY: build
build: build-frontend build-backend

.PHONY: clean-frontend
clean-frontend:
	rm -rf src/frontend/node_modules
	rm -f src/frontend/.eslintcache
	rm -rf src/frontend/dist
	rm -rf dist/dist
	rm -rf dist/frontend

.PHONY: clean-backend
clean-backend:
	rm -rf src/backend/embed
	rm -f src/backend/dlp-ui
	rm -rf dist/backend

.PHONY: clean
clean: clean-frontend clean-backend
	rm -rf dist