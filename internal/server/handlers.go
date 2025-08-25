// Package server provides the HTTP server setup, including
// route registration, handlers, and static file serving.
package server

import (
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
	mux.HandleFunc("POST /submit", handlerSubmit)
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
		http.Error(w, "Erro ao processar formulário", http.StatusBadRequest)
		return
	}

	prompt := r.PostForm.Get("prompt")
	fmt.Fprintf(w, "Você escreveu: %s", prompt)
}
