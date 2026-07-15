package outputs

import (
	"encoding/json"
	"strconv"
	"strings"
)

func ReadObject(content []byte) (map[string]any, error) {
	var object map[string]any
	err := json.Unmarshal([]byte(content), &object)
	if err != nil {
		return nil, err
	}

	return object, err
}

func ReadAria2(content string) (string, int, int, int) {
	if strings.HasPrefix(content, "[#") && strings.HasSuffix(content, "]") {
		content = strings.Trim(content, "[]")

		fields := strings.Fields(content)

		progress, _, _ := strings.Cut(fields[1], "(")
		current, total, _ := strings.Cut(progress, "/")

		_, speed, _ := strings.Cut(fields[3], ":")

		current = strings.TrimSuffix(current, "B")
		total = strings.TrimSuffix(total, "B")
		speed = strings.TrimSuffix(speed, "B")

		c, err := strconv.Atoi(current)
		if err != nil {
			c = -1
		}

		t, err := strconv.Atoi(total)
		if err != nil {
			t = -1
		}

		s, err := strconv.Atoi(speed)
		if err != nil {
			s = -1
		}

		return "", c, t, s
	}

	title, ok := strings.CutPrefix(content, "TITLE: ")
	if ok {
		return title, -1, -1, -1
	}

	return "", -1, -1, -1
}
