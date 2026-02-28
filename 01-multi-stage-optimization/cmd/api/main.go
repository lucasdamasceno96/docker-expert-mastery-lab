package main

import (
	"database/sql"
	"log"
	"net/http"

	"minimalrestapi/internal/handler"

	"minimalrestapi/internal/repository"

	_ "modernc.org/sqlite"
)

func main() {
	// 1. Setup Infra
	db, err := sql.Open("sqlite", "/data/app.db")
	if err != nil {
		log.Fatal(err)
	}

	// 2. Setup Repository
	repo := repository.NewLogRepository(db)
	if err := repo.Bootstrap(); err != nil {
		log.Fatal(err)
	}

	// 3. Setup Handler
	h := &handler.LogHandler{Repo: repo}

	// 4. Routes
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", h.HealthCheck)
	mux.HandleFunc("POST /logs", h.CreateLog)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
