package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadTitle(t *testing.T) {
	tests := []struct {
		Name string
		Case string
		Want string
	}{
		{"Normal", "TITLE: test", "test"},
		{"Space", "TITLE: te st", "te st"},
		{"Seperator", "TITLE: te: st", "te: st"},
		{"Stdout", "test test test", ""},
		{"Stderr", "ERROR: test", ""},
		{"Empty", "", ""},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := ReadTitle(test.Case)
			require.Equal(t, test.Want, got)
		})
	}
}

func TestReadProgress(t *testing.T) {
	tests := []struct {
		Name  string
		Case  string
		Wants []int
	}{
		{"Normal", "[#normal 1B/2B(50%) CN:1 DL:1B ETA:1s]", []int{1, 2, 1}},
		{"ETA", "[#noneta 1B/2B(50%) CN:1 DL:0B]", []int{1, 2, 0}},
		{"Form", "[test test test test test]", []int{-1, -1, -1}},
		{"Stdout", "test test test test test", []int{-1, -1, -1}},
		{"Empty", "", []int{-1, -1, -1}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got1, got2, got3 := ReadAria2(test.Case)
			gots := []int{got1, got2, got3}
			require.Equal(t, test.Wants, gots)
		})
	}
}
