package utils

import "strings"

const (
	Reset     = "\033[0m"
	Strong    = "\033[1m"
	Fade      = "\033[2m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
	Reverse   = "\033[7m"

	Black  = "\033[30m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Pink   = "\033[35m"
	Sky    = "\033[36m"
	White  = "\033[37m"
)

func Style(content string, styles ...string) string {
	return strings.Join(styles, "") + content + Reset
}
