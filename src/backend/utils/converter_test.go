package utils

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAtoi(t *testing.T) {
	cases := [4]string{
		"1",
		"2.",
		"3.3",
		"4.7",
	}
	wants := [4]int{
		1,
		2,
		3,
		4,
	}

	for index := range cases {
		t.Run(cases[index], func(t *testing.T) {
			got := Atoi(cases[index])
			require.Equal(t, wants[index], got)
		})
	}
}

func BenchmarkAtoi(b *testing.B) {
	cases := [4]string{
		"1",
		"2.",
		"3.2",
		"4.8",
	}

	for n := range b.N {
		Atoi(cases[n%len(cases)])
	}
}

func TestStom(t *testing.T) {
	cases := [2]sync.Map{
		{},
		{},
	}
	cases[0].Store("test", "test")
	wants := [2]map[string]any{
		{"test": "test"},
		{},
	}

	for index := range cases {
		t.Run(fmt.Sprint(wants[index]), func(t *testing.T) {
			m := Stom(&cases[index])
			require.Equal(t, wants[index], m)
		})
	}
}

func BenchmarkStom(b *testing.B) {
	cases := [2]sync.Map{
		{},
		{},
	}
	cases[0].Store("test", "test")

	for n := range b.N {
		Stom(&cases[n%len(cases)])
	}
}
