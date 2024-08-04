package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths of the slices are not equal:  actual: '%v' vs expected : '%v'", len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			if actual[i] != cs.expected[i] {
				t.Errorf("actual: '%v' does not equal expected: '%v'", actual[i], cs.expected[i])
			}
		}
	}
}
