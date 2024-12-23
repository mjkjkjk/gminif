package main

import (
	"testing"
)

func TestMinifyJSON(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    string
		shouldError bool
	}{
		{
			name: "simple object",
			input: `{
				"name": "John",
				"age": 30,
				"city": "New York"
			}`,
			expected: `{"age":30,"city":"New York","name":"John"}`,
		},
		{
			name: "nested object",
			input: `{
				"person": {
					"name": "John",
					"age": 30
				},
				"active": true
			}`,
			expected: `{"active":true,"person":{"age":30,"name":"John"}}`,
		},
		{
			name:        "invalid JSON",
			input:       `{"name": "John", "age": }`,
			expected:    "",
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := minifyJSON(tc.input)
			if tc.shouldError {
				if err == nil {
					t.Error("expected error but got none")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, result)
			}
		})
	}
}
