package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello from gminif!")
}

func minifyJSON(input string) (string, error) {
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
