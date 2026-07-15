package outputs

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestReadObject(t *testing.T) {
	tests := []struct {
		Name string
		Case []byte
		Want map[string]any
	}{
		{
			"Object",
			[]byte("{\"test\": \"test\"}"),
			map[string]any{"test": "test"},
		},
		{
			"Array",
			[]byte("[{\"test\": \"test\"}]"),
			nil,
		},
		{
			"Char",
			[]byte("{\"test\": \"t\\\\nes\\\\tt\"}"),
			map[string]any{"test": "t\\nes\\tt"},
		},
		{
			"Invalid",
			[]byte("[{\"test\": }]"),
			nil,
		},
		{
			"Log",
			[]byte("test"),
			nil,
		},
		{
			"Empty",
			[]byte("\n \n"),
			nil,
		},
		{
			"None",
			[]byte(""),
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got, _ := ReadObject(test.Case)

			diff := cmp.Diff(test.Want, got)
			if diff != "" {
				t.Fatalf("\n--- expected\n+++ actual\n%s", diff)
			}
		})
	}
}

func TestReadAria2(t *testing.T) {
	t.Run("Title", func(t *testing.T) {
		tests := []struct {
			Name string
			Case string
			Want string
		}{
			{
				"Normal",
				"TITLE: test",
				"test",
			},
			{
				"Space",
				"TITLE: te st",
				"te st",
			},
			{
				"Char",
				"TITLE: te\tst",
				"te\tst",
			},
			{
				"Form",
				"FORM: test",
				"",
			},
			{
				"Prefix",
				"TITLE: teTITLE: st",
				"teTITLE: st",
			},
			{
				"Log",
				"test test test test test",
				"",
			},
			{
				"Progress",
				"[#progre 1B/2B(50%) CN:1 DL:1B ETA:1s]",
				"",
			},
			{
				"Empty",
				"\n \n",
				"",
			},
			{
				"None",
				"",
				"",
			},
		}

		for _, test := range tests {
			t.Run(test.Name, func(t *testing.T) {
				got, _, _, _ := ReadAria2(test.Case)
				require.Equal(t, test.Want, got)
			})
		}
	})

	t.Run("Progress", func(t *testing.T) {
		tests := []struct {
			Name string
			Case string
			Want []int
		}{
			{
				"Normal",
				"[#normal 1B/2B(50%) CN:1 DL:1B ETA:1s]",
				[]int{1, 2, 1},
			},
			{
				"ETA",
				"[#etaeta 1B/2B(50%) CN:1 DL:0B]",
				[]int{1, 2, 0},
			},
			{
				"Form",
				"[#test test test test test]",
				[]int{-1, -1, -1},
			},
			{
				"Log",
				"test test test test test",
				[]int{-1, -1, -1},
			},
			{
				"Title",
				"TITLE: test",
				[]int{-1, -1, -1},
			},
			{
				"Empty",
				"\n \n",
				[]int{-1, -1, -1},
			},
			{
				"None",
				"",
				[]int{-1, -1, -1},
			},
		}

		for _, test := range tests {
			t.Run(test.Name, func(t *testing.T) {
				_, got1, got2, got3 := ReadAria2(test.Case)
				got := []int{got1, got2, got3}

				diff := cmp.Diff(test.Want, got)
				if diff != "" {
					t.Fatalf("\n--- expected\n+++ actual\n%s", diff)
				}
			})
		}
	})
}
