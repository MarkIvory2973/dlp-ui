package ytdlp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewParser(t *testing.T) {
	logger := setup(t)

	tests := []struct {
		Name string
		Case string
	}{
		{"Unsupported", "https://TestNewParser/Unsupported"},
		{"Invalid", "TestNewParser/Invalid"},
		{"Empty", ""},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			parseds := []Parsed{{URL: test.Case}}

			parser, err := NewParser("", test.Case, parseds)
			require.NoError(t, err)

			parser(logger)

			require.NotEmpty(t, parseds)

			for _, parsed := range parseds {
				require.NotEmpty(t, parsed.Errors)
			}
		})
	}
}
