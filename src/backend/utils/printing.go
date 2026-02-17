package utils

import "strings"

const Reset = "\033[0m"

const Highlight = "\033[1m"
const Fade = "\033[2m"

const Green = "\033[32m"
const SkyBlue = "\033[36m"

func UseStyles(a string, style ...string) string {
	return strings.Join(style, "") + a + Reset
}
