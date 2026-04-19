package utils

import "strconv"

func Atoi(content string, fallback int) int {
	number, err := strconv.Atoi(content)
	if err != nil {
		return fallback
	}

	return number
}
