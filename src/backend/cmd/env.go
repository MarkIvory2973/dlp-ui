package cmd

import (
	"os"
	"slices"
	"strings"
)

var supportedBrowsers = []string{
	"brave",
	"chrome",
	"chromium",
	"edge",
	"firefox",
	"opera",
	"safari",
	"vivaldi",
	"whale",
}

func GetDebug() bool {
	debug := os.Getenv("DEBUG")
	debug = strings.TrimSpace(debug)
	debug = strings.ToLower(debug)

	if debug != "true" {
		return false
	}

	return true
}

func GetBrowser() string {
	browser := os.Getenv("BROWSER")
	browser = strings.TrimSpace(browser)
	browser = strings.ToLower(browser)

	if !slices.Contains(supportedBrowsers, browser) {
		return ""
	}

	return browser
}
