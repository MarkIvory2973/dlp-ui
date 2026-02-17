package utils

import (
	"bufio"
	"encoding/json"
	"io"
	"strings"
)

func ReadLines(reader *bufio.Reader) ([]string, error) {
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		line = strings.Trim(line, "\n")
		line = strings.TrimSpace(line)

		if line != "" {
			lines = append(lines, line)
		}

		if err != nil {
			if err == io.EOF {
				break
			}

			return lines, err
		}
	}

	return lines, nil
}

func ReadJsons(reader *bufio.Reader) ([]map[string]any, error) {
	lines, err := ReadLines(reader)
	if err != nil {
		return nil, err
	}

	var objects []map[string]any
	for _, line := range lines {
		var object map[string]any
		err := json.Unmarshal([]byte(line), &object)
		if err != nil {
			continue
		}

		objects = append(objects, object)
	}

	return objects, err
}

func ReadTitle(line string) string {
	if len(line) < 7 {
		return ""
	}

	blocks := strings.Split(line, ": ")

	header := blocks[0]
	if header != "TITLE" {
		return ""
	}

	title := blocks[1]

	return title
}

func ReadProgress(line string) (int, int, int) {
	if len(line) < 3 {
		return -1, -1, -1
	}

	if line[0] != '[' && line[len(line)-1] != ']' {
		return -1, -1, -1
	}

	line = strings.Trim(line, "[")
	line = strings.Trim(line, "]")

	blocks := strings.Split(line, " ")
	if len(blocks) < 4 && 5 < len(blocks) {
		return -1, -1, -1
	}

	blocks2 := strings.Split(blocks[1], "(")
	if len(blocks2) != 2 {
		return -1, -1, -1
	}

	blocks2_1 := strings.Split(blocks2[0], "/")
	if len(blocks2_1) != 2 {
		return -1, -1, -1
	}

	block_current := strings.Trim(blocks2_1[0], "B")
	block_total := strings.Trim(blocks2_1[1], "B")

	blocks4 := strings.Split(blocks[3], ":")
	if len(blocks4) != 2 {
		return -1, -1, -1
	}

	block_speed := strings.Trim(blocks4[1], "B")

	current := Atoi(block_current)
	total := Atoi(block_total)
	speed := Atoi(block_speed)

	return current, total, speed
}
