package main

import (
	"fmt"
	"log"

	"github.com/fematarazzo/desculpaai/internal/server"
)

func main() {
	srv := server.New()
	fmt.Println("Server rodando em http://localhost:8080 ...")
	log.Fatal(srv.ListenAndServe(":8080"))
}
