package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

func ScanLineFunc(pipe io.ReadCloser, handler func(string)) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		content := scanner.Text()
		content = strings.TrimSpace(content)
		if content == "" {
			continue
		}

		handler(content)
	}
}

func ScanJsonFunc(pipe io.ReadCloser, handler func(map[string]any, error)) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		content := scanner.Bytes()
		content = bytes.TrimSpace(content)
		if content == nil {
			continue
		}

		var object map[string]any
		err := json.Unmarshal(content, &object)

		handler(object, err)
	}
}
