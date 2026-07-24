GO := go
GOFLAGS := -trimpath -ldflags="-s -w"
CGO_ENABLED := 0
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

UPX := upx
UPXFLAGS := --best --lzma

# Install dependencies
.install: go.mod go.sum
	$(GO) mod download

	touch $@

.PHONY: install
install: .install

# Test units
embed/frontend/test:
	mkdir -p embed/frontend
	touch embed/frontend/test

.PHONY: test
test: embed/frontend/test
	$(GO) test ./...

# Build binaries
output: .install $(shell find . -type f -iname "*.go")
	$(GO) build $(GOFLAGS) -o output
	-$(UPX) $(UPXFLAGS) output

.PHONY: build
build: output

# Clean files
.PHONY: clean
clean:
	$(RM) embed/frontend/test
	$(RM) .install
	$(RM) output