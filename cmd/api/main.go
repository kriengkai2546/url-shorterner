package main

import (
	"log"
	"net/http"
	"os"
	"urlshortener/internal/auth"
	"urlshortener/pkg/database"

	"github.com/joho/godotenv"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file, using system env")
	}

	// connect DB
	db := database.Connect()
	defer db.Close()

	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := *auth.NewHandler(authService)



	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	mux.HandleFunc("/auth/register", authHandler.Register)
	mux.HandleFunc("/auth/login", authHandler.Login)

	// start server
	port := os.Getenv("PORT")
	log.Printf("Server running on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
