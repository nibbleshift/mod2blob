package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_parseFunction(t *testing.T) {
	var tests = []struct {
		definition string
		expected   *Function
		err        error
	}{
		{
			definition: "func test(test int)",
			expected:   nil,
			err:        nil,
		},
		{
			definition: "func Echo(test string, x float64) []string",
			expected: &Function{
				Name: "Echo",
				Args: []Arg{
					{
						Name: "test",
						Type: "string",
					},
					{
						Name: "x",
						Type: "float64",
					},
				},
				Return: []Arg{
					{
						Type: "[]string",
					},
				},
			},
			err: nil,
		},
		{
			definition: "func testtest int)",
			expected:   nil,
			err:        ErrInvalidFunction,
		},
		{
			definition: "func Echo2(test, x float64) ([]string, error)",
			expected: &Function{
				Name: "Echo2",
				Args: []Arg{
					{
						Name: "test",
						Type: "float64",
					},
					{
						Name: "x",
						Type: "float64",
					},
				},
				Return: []Arg{
					{
						Type: "[]string",
					},
					{
						Type: "error",
					},
				},
			},
			err: nil,
		},
		{
			definition: "func testtest int)",
			expected:   nil,
			err:        ErrInvalidFunction,
		},
	}

	for _, tt := range tests {
		t.Run(tt.definition, func(t *testing.T) {
			parsed, err := parseFunction(tt.definition)
			assert.Equal(t, err, tt.err)
			assert.DeepEqual(t, parsed, tt.expected)
		})
	}
}
