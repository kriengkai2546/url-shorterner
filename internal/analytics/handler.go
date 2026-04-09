package analytics

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	log.Printf("path: %s", r.URL.Path)
	idStr := strings.TrimPrefix(r.URL.Path, "/analytics/")
	log.Printf("idStr: '%s'", idStr) 
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	count, err := h.repo.GetClickCount(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"clicks": count})
}
