package outputs

import (
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestScanTextFunc(t *testing.T) {
	tests := []struct {
		Name string
		Case string
		Want []string
	}{
		{
			"LF",
			"test\ntest",
			[]string{"test", "test"},
		},
		{
			"CRLF",
			"test\r\ntest",
			[]string{"test", "test"},
		},
		{
			"Spaces",
			" test \n test ",
			[]string{"test", "test"},
		},
		{
			"Tabs",
			"\ttest\t\n\ttest\t",
			[]string{"test", "test"},
		},
		{
			"Trailings",
			"test\ntest\n",
			[]string{"test", "test"},
		},
		{
			"Series",
			"test\n \ntest",
			[]string{"test", "test"},
		},
		{
			"Empty",
			"\n \n",
			[]string{},
		},
		{
			"None",
			"",
			[]string{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			pipe := io.NopCloser(strings.NewReader(test.Case))

			got := []string{}

			ScanTextFunc(pipe, func(content string) {
				got = append(got, content)
			})

			diff := cmp.Diff(test.Want, got)
			if diff != "" {
				t.Fatalf("\n--- expected\n+++ actual\n%s", diff)
			}
		})
	}
}

func TestScanBytesFunc(t *testing.T) {
	tests := []struct {
		Name string
		Case string
		Want [][]byte
	}{
		{
			"LF",
			"test\ntest",
			[][]byte{[]byte("test"), []byte("test")},
		},
		{
			"CRLF",
			"test\r\ntest",
			[][]byte{[]byte("test"), []byte("test")},
		},
		{
			"Spaces",
			" test \n test ",
			[][]byte{[]byte("test"), []byte("test")},
		},
		{
			"Tabs",
			"\ttest\t\n\ttest\t",
			[][]byte{[]byte("test"), []byte("test")},
		},
		{
			"Trailings",
			"test\ntest\n",
			[][]byte{[]byte("test"), []byte("test")},
		},
		{
			"Series",
			"test\n \ntest",
			[][]byte{[]byte("test"), []byte("test")},
		},
		{
			"Empty",
			"\n \n",
			[][]byte{},
		},
		{
			"None",
			"",
			[][]byte{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			pipe := io.NopCloser(strings.NewReader(test.Case))

			got := [][]byte{}

			ScanBytesFunc(pipe, func(content []byte) {
				got = append(got, content)
			})

			diff := cmp.Diff(test.Want, got)
			if diff != "" {
				t.Fatalf("\n--- expected\n+++ actual\n%s", diff)
			}
		})
	}
}

func TestScanObjectFunc(t *testing.T) {
	tests := []struct {
		Name string
		Case string
		Want []map[string]any
	}{
		{
			"Objects",
			"{\"test\": \"test\"}" + "\n" + "{\"test\": \"test\"}",
			[]map[string]any{{"test": "test"}, {"test": "test"}},
		},
		{
			"Arrays",
			"{\"test\": \"test\"}" + "\n" + "[{\"test\": \"test\"}]",
			[]map[string]any{{"test": "test"}},
		},
		{
			"Chars",
			"{\"test\": \"test\"}" + "\n" + "{\"test\": \"te\\\\tst\"}" + "\n" + "{\"test\": \"te\\\\nst\"}",
			[]map[string]any{{"test": "test"}, {"test": "te\\tst"}, {"test": "te\\nst"}},
		},
		{
			"Invalids",
			"{\"test\": \"test\"}" + "\n" + "{\"test\": }" + "\n" + "{",
			[]map[string]any{{"test": "test"}},
		},
		{
			"Logs",
			"test" + "\n" + "{\"test\": \"test\"}" + "\n" + "test",
			[]map[string]any{{"test": "test"}},
		},
		{
			"Empty",
			"\n \n",
			[]map[string]any{},
		},
		{
			"None",
			"",
			[]map[string]any{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			pipe := io.NopCloser(strings.NewReader(test.Case))

			got := []map[string]any{}

			ScanObjectFunc(pipe, func(object map[string]any) {
				got = append(got, object)
			})

			diff := cmp.Diff(test.Want, got)
			if diff != "" {
				t.Fatalf("\n--- expected\n+++ actual\n%s", diff)
			}
		})
	}
}

func TestScanAria2Func(t *testing.T) {
	t.Run("Title", func(t *testing.T) {
		tests := []struct {
			Name string
			Case string
			Want []string
		}{
			{
				"Normal",
				"TITLE: test" + "\n" + "TITLE: test",
				[]string{"test", "test"},
			},
			{
				"Spaces",
				"TITLE: test" + "\n" + "TITLE: te st",
				[]string{"test", "te st"},
			},
			{
				"Chars",
				"TITLE: test" + "\n" + "TITLE: te\tst",
				[]string{"test", "te\tst"},
			},
			{
				"Forms",
				"TITLE: test" + "\n" + "FORM: test",
				[]string{"test"},
			},
			{
				"Prefixes",
				"TITLE: test" + "\n" + "TITLE: teTITLE: st",
				[]string{"test", "teTITLE: st"},
			},
			{
				"Logs",
				"test" + "\n" + "TITLE: test" + "\n" + "test",
				[]string{"test"},
			},
			{
				"Progresses",
				"[#progre 1B/2B(50%) CN:1 DL:1B ETA:1s]" + "\n" + "TITLE: test" + "\n" + "[#progre 1B/2B(50%) CN:1 DL:1B ETA:1s]",
				[]string{"test"},
			},
			{
				"Empty",
				"\n \n",
				[]string{},
			},
			{
				"None",
				"",
				[]string{},
			},
		}

		for _, test := range tests {
			t.Run(test.Name, func(t *testing.T) {
				pipe := io.NopCloser(strings.NewReader(test.Case))

				got := []string{}

				ScanAria2Func(pipe, func(title string, _ int, _ int, _ int) {
					if title == "" {
						return
					}

					got = append(got, title)
				})

				diff := cmp.Diff(test.Want, got)
				if diff != "" {
					t.Fatalf("\n--- expected\n+++ actual\n%s", diff)
				}
			})
		}
	})

	t.Run("Progress", func(t *testing.T) {
		tests := []struct {
			Name string
			Case string
			Want [][]int
		}{
			{
				"Normal",
				"[#normal 1B/2B(50%) CN:1 DL:1B ETA:1s]" + "\n" + "[#normal 1B/2B(50%) CN:1 DL:1B ETA:1s]",
				[][]int{{1, 2, 1}, {1, 2, 1}},
			},
			{
				"ETAs",
				"[#etaset 1B/2B(50%) CN:1 DL:1B ETA:1s]" + "\n" + "[#etaset 1B/2B(50%) CN:1 DL:0B]",
				[][]int{{1, 2, 1}, {1, 2, 0}},
			},
			{
				"Forms",
				"[#formsf 1B/2B(50%) CN:1 DL:1B ETA:1s]" + "\n" + "[#formsf test test test test]",
				[][]int{{1, 2, 1}, {-1, -1, -1}},
			},
			{
				"Logs",
				"test" + "\n" + "[#logslo 1B/2B(50%) CN:1 DL:1B ETA:1s]" + "\n" + "test",
				[][]int{{-1, -1, -1}, {1, 2, 1}, {-1, -1, -1}},
			},
			{
				"Titles",
				"TITLE: test" + "\n" + "[#titles 1B/2B(50%) CN:1 DL:1B ETA:1s]" + "\n" + "TITLE: test",
				[][]int{{-1, -1, -1}, {1, 2, 1}, {-1, -1, -1}},
			},
			{
				"Empty",
				"\n \n",
				[][]int{},
			},
			{
				"None",
				"",
				[][]int{},
			},
		}

		for _, test := range tests {
			t.Run(test.Name, func(t *testing.T) {
				pipe := io.NopCloser(strings.NewReader(test.Case))

				got := [][]int{}

				ScanAria2Func(pipe, func(_ string, current int, total int, speed int) {
					got = append(got, []int{current, total, speed})
				})

				diff := cmp.Diff(test.Want, got)
				if diff != "" {
					t.Fatalf("\n--- expected\n+++ actual\n%s", diff)
				}
			})
		}
	})
}
