package url

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"urlshortener/internal/auth"
)

type Handler struct {
	service *Service
}

func NewHandler(Service *Service) *Handler {
	return &Handler{service: Service}
}

func (h *Handler) CreateURL(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.UserIDKey).(int)

	var req CreateURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	result, err := h.service.CreateURL(userID, req.LongURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	longURL, err := h.service.GetOriginalURL(code)
	if err != nil {
		http.Error(w, "url not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL,http.StatusMovedPermanently)
}

func (h *Handler) GetUserURLs(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.UserIDKey).(int)

	urls, err := h.service.GetUserURLs(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInsufficientStorage)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}

func (h *Handler) DeleteURL(w http.ResponseWriter, r *http.Request) {
	userID := r. Context().Value(auth.UserIDKey).(int)
	idStr := strings.TrimPrefix(r.URL.Path, "/urls/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteURL(id, userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
