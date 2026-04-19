package utils

import (
	"slices"
	"strings"

	"github.com/ua-parser/uap-go/uaparser"
)

var supportedBrowsers = []string{
	"chrome",
	"edge",
	"firefox",
	"opera",
	"safari",
}

func ParseBrowser(ua string) (string, error) {
	parser, err := uaparser.New()
	if err != nil {
		return "", err
	}

	userAgent := parser.ParseUserAgent(ua)
	currentBrowser := strings.ToLower(userAgent.Family)
	if !slices.Contains(supportedBrowsers, currentBrowser) {
		return "", nil
	}

	return currentBrowser, nil
}
