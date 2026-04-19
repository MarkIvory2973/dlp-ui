package utils

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScanLineFunc(t *testing.T) {
	tests := []struct {
		Name string
		Case string
		Want []string
	}{
		{"1", "test\ntest", []string{"test", "test"}},
		{"2", "test\ntest\n", []string{"test", "test"}},
		{"3", "test \ntest ", []string{"test", "test"}},
		{"4", "test \ntest \n", []string{"test", "test"}},
		{"5", " test\n test", []string{"test", "test"}},
		{"6", " test\n test\n", []string{"test", "test"}},
		{"7", " test \n test ", []string{"test", "test"}},
		{"8", " test \n test \n", []string{"test", "test"}},
		{"9", "test\n", []string{"test"}},
		{"10", "test \n", []string{"test"}},
		{"11", " test\n", []string{"test"}},
		{"12", " test \n", []string{"test"}},
		{"13", "\n", nil},
		{"14", "\n\n", nil},
		{"15", " \n ", nil},
		{"16", " \n \n", nil},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			pipe := io.NopCloser(strings.NewReader(test.Case))

			var got []string

			ScanLineFunc(pipe, func(content string) {
				got = append(got, content)
			})

			require.Equal(t, test.Want, got)
		})
	}
}

func TestScanJsonFunc(t *testing.T) {
	tests := []struct {
		Name  string
		Case  string
		Wants []string
	}{
		{"1", `{"test": "test"}` + "\n" + `{"test": "test"}`, []string{`{"test": "test"}`, `{"test": "test"}`}},
		{"2", `{"test": "test\ntest"}` + "\n" + `{"test": "test\ntest"}`, []string{`{"test": "test\ntest"}`, `{"test": "test\ntest"}`}},
		{"3", `{"test": "test"}`, []string{`{"test": "test"}`}},
		{"4", `{"test": "test\ntest"}`, []string{`{"test": "test\ntest"}`}},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			pipe := io.NopCloser(strings.NewReader(test.Case))

			var gots []string

			ScanJsonFunc(pipe, func(object map[string]any, err error) {
				content, err := json.Marshal(object)
				require.NoError(t, err)

				gots = append(gots, string(content))
			})

			for index := range test.Wants {
				require.JSONEq(t, test.Wants[index], gots[index])
			}
		})
	}
}
