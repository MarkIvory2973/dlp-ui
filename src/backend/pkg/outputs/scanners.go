package outputs

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

func ScanTextFunc(pipe io.ReadCloser, handler func(string)) {
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

func ScanBytesFunc(pipe io.ReadCloser, handler func([]byte)) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		content := scanner.Bytes()
		content = bytes.TrimSpace(content)
		if content == nil {
			continue
		}

		handler(content)
	}
}

func ScanObjectFunc(pipe io.ReadCloser, handler func(map[string]any)) {
	ScanBytesFunc(pipe, func(content []byte) {
		object, err := ReadObject(content)
		if err != nil {
			return
		}

		handler(object)
	})
}

func ScanAria2Func(pipe io.ReadCloser, handler func(title string, current int, total int, speed int)) {
	ScanTextFunc(pipe, func(content string) {
		title, current, total, speed := ReadAria2(content)
		handler(title, current, total, speed)
	})
}
