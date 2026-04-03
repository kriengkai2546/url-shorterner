package main

import (
	"log"
	"net/http"
	"os"
	"urlshortener/pkg/database"

	"github.com/joho/godotenv"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file, using system env")
	}

	// เพิ่มบรรทัดนี้
	log.Printf("DB_HOST=%s DB_NAME=%s", os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// connect DB
	db := database.Connect()
	defer db.Close()

	// start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	log.Printf("Server running on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
