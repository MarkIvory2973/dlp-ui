package utils

import (
	"strings"
)

func ReadTitle(content string) string {
	if !strings.HasPrefix(content, "TITLE: ") {
		return ""
	}

	title := strings.TrimPrefix(content, "TITLE: ")

	return title
}

func ReadAria2(content string) (int, int, int) {
	if !strings.HasPrefix(content, "[#") {
		return -1, -1, -1
	} else if !strings.HasSuffix(content, "]") {
		return -1, -1, -1
	}

	content = strings.Trim(content, "[]")

	fields := strings.Fields(content)

	progress, _, _ := strings.Cut(fields[1], "(")
	current, total, _ := strings.Cut(progress, "/")

	_, speed, _ := strings.Cut(fields[3], ":")

	current = strings.TrimSuffix(current, "B")
	total = strings.TrimSuffix(total, "B")
	speed = strings.TrimSuffix(speed, "B")

	return Atoi(current, -1), Atoi(total, -1), Atoi(speed, -1)
}
