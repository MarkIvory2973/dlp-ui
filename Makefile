include mk/output.mk
include mk/nfpm.mk

# Install dependencies
.PHONY: install-frontend
install-frontend:
	$(MAKE) -C src/frontend install

.PHONY: install-backend
install-backend:
	$(MAKE) -C src/backend install

.PHONY: install
install: install-frontend install-backend

# Test units
.PHONY: test-frontend
test-frontend:
	$(MAKE) -C src/frontend test

.PHONY: test-backend
test-backend:
	$(MAKE) -C src/backend test

.PHONY: test
test: test-frontend test-backend

# Build frontend and backend
dist/frontend/:
	$(MAKE) -C src/frontend build
	mkdir -p dist
	$(RM) -r dist/frontend
	mv src/frontend/frontend dist

.PHONY: build-frontend
build-frontend: dist/frontend/

src/backend/embed/frontend/: dist/frontend/
	mkdir -p src/backend/embed
	cp -r dist/frontend src/backend/embed

dist/backend/output: src/backend/embed/frontend/
	$(MAKE) -C src/backend build
	mkdir -p dist/backend
	mv src/backend/output dist/backend

.PHONY: build-backend
build-backend: dist/backend/output
ifeq ($(OS),linux)
	$(NFPM) pkg --packager deb $(NFPMFLAGS)
	$(NFPM) pkg --packager rpm $(NFPMFLAGS)
endif
	mv dist/backend/output dist/backend/$(OUTPUT)

.PHONY: build
build: build-backend

# Clean files
.PHONY: clean-frontend
clean-frontend:
	$(MAKE) -C src/frontend clean
	$(RM) -r dist/frontend

.PHONY: clean-backend
clean-backend:
	$(RM) -r src/backend/embed
	$(MAKE) -C src/backend clean
	$(RM) -r dist/backend

.PHONY: clean
clean: clean-frontend clean-backend
	$(RM) -r dist