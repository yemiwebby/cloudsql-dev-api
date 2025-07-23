package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	log.Printf("Connecting to database with config: Host=%s, Port=%s, User=%s, DB=%s", dbHost, dbPort, dbUser, dbName)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)

	log.Printf("Connection string: postgres://%s:***@%s:%s/%s?sslmode=disable", dbUser, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to create DB connection pool: %v", err)
	}

	defer db.Close()

	// Test the database connection
	log.Printf("Testing database connection...")
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}
	log.Printf("Database connection successful!")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			log.Printf("Health check failed: %v", err)
			http.Error(w, "DB unreachable", http.StatusServiceUnavailable)
			return
		}
		w.Write([]byte("OK"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("API running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
