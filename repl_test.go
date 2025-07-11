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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},

		{
			input:    "  how art   thou  ",
			expected: []string{"how", "art", "thou"},
		},

		{
			input:    "   leading spaces",
			expected: []string{"leading", "spaces"},
		},
		{
			input:    "trailing spaces   ",
			expected: []string{"trailing", "spaces"},
		},
		{
			input:    " multiple   spaces   between  words ",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		if len(actual) != len(c.expected) {
			t.Errorf("Actual Word Count less than Expected Word Count")
		}
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word == expectedWord {
				continue
			} else {
				t.Errorf("Expected Word = %s. Actual Word = %s", expectedWord, word)
			}
		}
	}
}
