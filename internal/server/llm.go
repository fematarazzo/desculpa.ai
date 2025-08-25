package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

var ollamaURL = os.Getenv("OLLAMA_URL")

func callOllama(prompt string) (string, error) {
	body := map[string]string{
		"model":  "gemma3:1b",
		"prompt": prompt,
	}
	b, _ := json.Marshal(body)

	resp, err := http.Post(ollamaURL+"/api/generate", "application/json", bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var parsed map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return "", err
	}

	if out, ok := parsed["response"].(string); ok {
		return out, nil
	}
	return "sem resposta", nil
}
