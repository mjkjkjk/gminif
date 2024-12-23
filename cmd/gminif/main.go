package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello from gminif!")
}

func validateJSON(input string) error {
	var js json.RawMessage
	if err := json.Unmarshal([]byte(input), &js); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}
	return nil
}

func minifyJSON(input string) (string, error) {
	if err := validateJSON(input); err != nil {
		return "", err
	}

	var parsed interface{}
	if err := json.Unmarshal([]byte(input), &parsed); err != nil {
		return "", err
	}

	minified, err := json.Marshal(parsed)
	if err != nil {
		return "", err
	}

	return string(minified), nil
}
