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
			name: "complex object",
			input: `{
  "product": "Live JSON generator",
  "version": 3.1,
  "releaseDate": "2014-06-25T00:00:00.000Z",
  "demo": true,
  "person": {
    "id": 12345,
    "name": "John Doe",
    "phones": {
      "home": "800-123-4567",
      "mobile": "877-123-1234"
    },
    "email": [
      "jd@example.com",
      "jd@example.org"
    ],
    "dateOfBirth": "1980-01-02T00:00:00.000Z",
    "registered": true,
    "emergencyContacts": [
      {
        "name": "Jane Doe",
        "phone": "888-555-1212",
        "relationship": "spouse"
      },
      {
        "name": "Justin Doe",
        "phone": "877-123-1212",
        "relationship": "parent"
      }
    ]
  }
}`,
			expected: `{"demo":true,"person":{"dateOfBirth":"1980-01-02T00:00:00.000Z","email":["jd@example.com","jd@example.org"],"emergencyContacts":[{"name":"Jane Doe","phone":"888-555-1212","relationship":"spouse"},{"name":"Justin Doe","phone":"877-123-1212","relationship":"parent"}],"id":12345,"name":"John Doe","phones":{"home":"800-123-4567","mobile":"877-123-1234"},"registered":true},"product":"Live JSON generator","releaseDate":"2014-06-25T00:00:00.000Z","version":3.1}`,
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
