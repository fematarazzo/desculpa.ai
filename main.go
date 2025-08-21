package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerMain(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello, world!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerMain)

	fmt.Println("Turning server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
