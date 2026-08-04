GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

ifeq ($(GOOS),windows)
OUTPUT := dlp-ui_$(GOOS)_$(GOARCH).exe
else
OUTPUT := dlp-ui_$(GOOS)_$(GOARCH)
endif