package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MinifyRequest struct {
	Input string `json:"input"`
}

type MinifyResponse struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func handleMinify(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req MinifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result, err := minifyJSON(req.Input)

	w.Header().Set("Content-Type", "application/json")

	response := MinifyResponse{}
	if err != nil {
		response.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response.Result = result
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/minify", handleMinify)

	fs := http.FileServer(http.Dir("web/dist"))
	http.Handle("/", fs)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
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
