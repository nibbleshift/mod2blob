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
			expected: &Function{
				Name: "test",
				Args: []Arg{
					{
						Name: "test",
						Type: "int",
					},
				},
			},
			err: nil,
		},
		{
			definition: "func test(test []string, two float64, four map[string]string)",
			expected: &Function{
				Name: "test",
				Args: []Arg{
					{
						Name: "test",
						Type: "[]string",
					},
					{
						Name: "two",
						Type: "float64",
					},
					{
						Name: "four",
						Type: "map[string]string",
					},
				},
			},
			err: nil,
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
			definition: "func NoArgs() ([]string, error)",
			expected: &Function{
				Name: "NoArgs",
				Args: nil,
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
			definition: "func Echo2(test, one, two, three, x float64)",
			expected: &Function{
				Name: "Echo2",
				Args: []Arg{
					{
						Name: "test",
						Type: "float64",
					},
					{
						Name: "one",
						Type: "float64",
					},
					{
						Name: "two",
						Type: "float64",
					},
					{
						Name: "three",
						Type: "float64",
					},
					{
						Name: "x",
						Type: "float64",
					},
				},
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.definition, func(t *testing.T) {
			actual, err := parseFunction(tt.definition)
			assert.Equal(t, err, tt.err)
			assert.DeepEqual(t, actual, tt.expected)
		})
	}
}

func Test_parseReturn(t *testing.T) {
	var tests = []struct {
		definition string
		expected   *Arg
		err        error
	}{
		{
			definition: "float64",
			expected: &Arg{
				Type: "float64",
			},
			err: nil,
		},
		{
			definition: "[]string",
			expected: &Arg{
				Type: "[]string",
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.definition, func(t *testing.T) {
			actual, err := parseReturn(tt.definition)
			assert.Equal(t, err, tt.err)
			assert.DeepEqual(t, actual, tt.expected)
		})
	}
}

func Test_parseReturnArguments(t *testing.T) {
	var tests = []struct {
		definition string
		expected   []Arg
		err        error
	}{
		{
			definition: "float64",
			expected: []Arg{
				{
					Type: "float64",
				},
			},
			err: nil,
		},
		{
			definition: "[]string",
			expected: []Arg{
				{
					Type: "[]string",
				},
			},
			err: nil,
		},
		{
			definition: "([]string, float64)",
			expected: []Arg{
				{
					Type: "[]string",
				},
				{
					Type: "float64",
				},
			},
			err: nil,
		},
		{
			definition: "([]string  float64)",
			expected:   nil,
			err:        ErrInvalidArguments,
		},
	}

	for _, tt := range tests {
		t.Run(tt.definition, func(t *testing.T) {
			actual, err := parseReturnArguments(tt.definition)
			assert.Equal(t, err, tt.err)
			assert.DeepEqual(t, actual, tt.expected)
		})
	}
}
