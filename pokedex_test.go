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
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		t.Logf("Testing With: %q", c.input)
		t.Logf("Expected: %v", c.expected)
		t.Logf("Actual: %v", actual)

		if len(actual) != len(c.expected) {
			t.Errorf("Length doesn't match: actual:%v expected:%v", len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Words don't match: actual:%v expected:%v", actual, c.expected)
			}
		}
	}
}
