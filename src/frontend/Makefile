NPM := npm

# Install dependencies
.install: package.json package-lock.json
	$(NPM) install

	touch $@

.PHONY: install
install: .install

# Test units
.PHONY: test
test: .install
	$(NPM) run lint

# Build frontend
frontend/: .install vite.config.js index.html $(shell find src -type f)
	$(NPM) run build

.PHONY: build
build: frontend/

# Clean files
.PHONY: clean
clean:
	$(RM) -r node_modules
	$(RM) .install
	$(RM) .eslintcache
	$(RM) -r frontend