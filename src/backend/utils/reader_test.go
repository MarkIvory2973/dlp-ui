package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadLines(t *testing.T) {
	cases := [8]string{
		"test\ntest",
		"test\ntest\n",
		"test \ntest ",
		"test \ntest \n",
		" test\n test",
		" test\n test\n",
		" test \n test ",
		" test \n test \n",
	}
	want := []string{"test", "test"}

	for index := range cases {
		r := strings.NewReader(cases[index])
		reader := bufio.NewReader(r)

		t.Run(cases[index], func(t *testing.T) {
			got, err := ReadLines(reader)
			require.NoError(t, err)

			require.Equal(t, want, got)
		})
	}
}

func BenchmarkReadLines(b *testing.B) {
	cases := [8]string{
		"test\ntest",
		"test\ntest\n",
		"test \ntest ",
		"test \ntest \n",
		" test\n test",
		" test\n test\n",
		" test \n test ",
		" test \n test \n",
	}

	for n := range b.N {
		r := strings.NewReader(cases[n%len(cases)])
		reader := bufio.NewReader(r)

		ReadLines(reader)
	}
}

func TestReadJsons(t *testing.T) {
	cases := [4]string{
		`{"test": "test"}`,
		`{"test": "test\ntest"}`,
		`{"test": "test"}` + "\n" + `{"test": "test"}`,
		`{"test": "test\ntest"}` + "\n" + `{"test": "test\ntest"}`,
	}
	wants := [4][]string{
		{`{"test": "test"}`},
		{`{"test": "test\ntest"}`},
		{`{"test": "test"}`, `{"test": "test"}`},
		{`{"test": "test\ntest"}`, `{"test": "test\ntest"}`},
	}

	for i := range cases {
		r := strings.NewReader(cases[i])
		reader := bufio.NewReader(r)

		t.Run(fmt.Sprint(cases[i]), func(t *testing.T) {
			objects, err := ReadJsons(reader)
			require.NoError(t, err)

			for j := range wants[i] {
				object, err := json.Marshal(objects[j])
				require.NoError(t, err)

				require.JSONEq(t, wants[i][j], string(object))
			}
		})
	}
}

func BenchmarkReadJsons(b *testing.B) {
	cases := [4]string{
		`{"test": "test"}`,
		`{"test": "test\ntest"}`,
		`{"test": "test"}` + "\n" + `{"test": "test"}`,
		`{"test": "test\ntest"}` + "\n" + `{"test": "test\ntest"}`,
	}

	for n := range b.N {
		r := strings.NewReader(cases[n%len(cases)])
		reader := bufio.NewReader(r)

		ReadJsons(reader)
	}
}

func TestReadTitle(t *testing.T) {
	cases := [5]string{
		"TITLE: test",
		"TITLE: te st",
		"test",
		"test test test",
		"ERROR: aria2c exited with code 33550337",
	}
	wants := [5]string{
		"test",
		"te st",
		"",
		"",
		"",
	}

	for index := range cases {
		t.Run(cases[index], func(t *testing.T) {
			got := ReadTitle(cases[index])
			require.Equal(t, wants[index], got)
		})
	}
}

func BenchmarkReadTitle(b *testing.B) {
	cases := [5]string{
		"TITLE: test",
		"TITLE: te st",
		"test",
		"test test test",
		"ERROR: aria2c exited with code 33550337",
	}

	for n := range b.N {
		ReadTitle(cases[n%len(cases)])
	}
}

func TestReadProgress(t *testing.T) {
	cases := [5]string{
		"[#test01 33550336B/33550337B(99%) CN:13 DL:33550336B ETA:33550336s]",
		"[#test02 0B/0B(0%) CN:0 DL:0B ETA:0s]",
		"[#test03 33550336B/33550337B(99%) CN:13 DL:33550336B]",
		"[#test04 0B/0B(0%) CN:0 DL:0B]",
		"ERROR: aria2c exited with code 33550337",
	}
	wants := [5][3]int{
		{33550336, 33550337, 33550336},
		{0, 0, 0},
		{33550336, 33550337, 33550336},
		{0, 0, 0},
		{-1, -1, -1},
	}

	for index := range cases {
		t.Run(cases[index], func(t *testing.T) {
			current, total, speed := ReadProgress(cases[index])

			got := [3]int{current, total, speed}
			require.Equal(t, wants[index], got)
		})
	}
}

func BenchmarkReadProgress(b *testing.B) {
	cases := [5]string{
		"[#test01 33550336B/33550337B(99%) CN:13 DL:33550336B ETA:33550336s]",
		"[#test02 0B/0B(0%) CN:0 DL:0B ETA:0s]",
		"[#test03 33550336B/33550337B(99%) CN:13 DL:33550336B]",
		"[#test04 0B/0B(0%) CN:0 DL:0B]",
		"ERROR: aria2c exited with code 33550337",
	}

	for n := range b.N {
		ReadProgress(cases[n%len(cases)])
	}
}
