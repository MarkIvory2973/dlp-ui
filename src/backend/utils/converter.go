package utils

import (
	"strconv"
	"strings"
	"sync"
)

func Atoi(s string) int {
	s = strings.Split(s, ".")[0]
	n, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}

	return n
}

func Stom(s *sync.Map) map[string]any {
	m := make(map[string]any)
	s.Range(func(k any, value any) bool {
		key, ok := k.(string)
		if !ok {
			return true
		}

		m[key] = value

		return true
	})

	return m
}
