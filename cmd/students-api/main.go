package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bedantchh/students-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students api"))
	})

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	fmt.Println("Server started")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}
}
