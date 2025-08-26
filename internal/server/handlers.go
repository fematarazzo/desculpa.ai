// Package server provides the HTTP server setup, including
// route registration, handlers, and static file serving.
package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func registerRoutes(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))
	mux.Handle("HEAD /static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("GET /", handlerMain)
	mux.HandleFunc("GET /contact", handlerContact)
	mux.HandleFunc("POST /submit", rateLimitMiddleware(handlerSubmit))
	mux.HandleFunc("POST /stream", rateLimitMiddleware(handlerStream))
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func handlerContact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/contact.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func handlerSubmit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error when processing form", http.StatusBadRequest)
		return
	}

	prompt := r.PostForm.Get("prompt")

	resp, err := callOllama(prompt)
	if err != nil {
		http.Error(w, "Ollama error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "gemma3:1b response: %s", resp)
}

func handlerStream(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error when processing form", http.StatusBadRequest)
		return
	}
	prompt := r.PostForm.Get("prompt")

	reqBody := map[string]any{
		"model":  "gemma3:1b",
		"prompt": prompt,
		"stream": true,
	}
	b, _ := json.Marshal(reqBody)

	resp, err := http.Post(ollamaURL+"/api/generate", "application/json", bytes.NewReader(b))
	if err != nil {
		http.Error(w, "Ollama error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(resp.Body)
	for {
		var msg map[string]any
		if err := decoder.Decode(&msg); err != nil {
			break
		}

		if token, ok := msg["response"].(string); ok {
			fmt.Fprintf(w, "%s", token)
			flusher.Flush()
		}

		if done, ok := msg["done"].(bool); ok && done {
			break
		}
	}
}
