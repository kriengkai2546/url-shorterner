package main

import (
	"log"
	"net/http"
	"os"

	"urlshortener/internal/analytics"
	"urlshortener/internal/auth"
	"urlshortener/internal/redirect"
	"urlshortener/internal/url"
	"urlshortener/pkg/cache"
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
	redisCache := cache.Connect()
	defer db.Close()

	// repositories
	authRepo := auth.NewRepository(db)
	urlRepo := url.NewRepository(db)
	analyticsRepo := analytics.NewRepository(db)

	// services
	authService := auth.NewService(authRepo)
	urlService := url.NewService(urlRepo)

	// handlers
	authHandler := auth.NewHandler(authService)
	urlHandler := url.NewHandler(urlService)
	redirectHandler := redirect.NewHandler(urlRepo, redisCache)
	analyticsHandler := analytics.NewHandler(analyticsRepo)

 mux := http.NewServeMux()
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status":"ok"}`))
    })
    mux.HandleFunc("/auth/register", authHandler.Register)
    mux.HandleFunc("/auth/login", authHandler.Login)
    mux.HandleFunc("/urls", auth.Middleware(urlHandler.CreateURL))
    mux.HandleFunc("/urls/", auth.Middleware(urlHandler.DeleteURL))
    mux.HandleFunc("/my-urls", auth.Middleware(urlHandler.GetUserURLs))
    mux.HandleFunc("/analytics/", auth.Middleware(analyticsHandler.GetStats))
    mux.HandleFunc("/", redirectHandler.Redirect)

    port := os.Getenv("PORT")
    log.Printf("Server running on :%s", port)
    http.ListenAndServe(":"+port, corsMiddleware(mux))
}


func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "https://url-shorterner-theta.vercel.app/")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}