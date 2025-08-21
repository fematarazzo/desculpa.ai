package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerMain)

	fmt.Println("Turning server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
