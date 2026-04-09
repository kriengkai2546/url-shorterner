package redirect

import (
	"net/http"
	"strings"
	"time"
	"urlshortener/internal/url"
	"urlshortener/pkg/cache"
)

type Handler struct {
	urlRepo *url.Repository
	cache *cache.Cache
}

func NewHandler(urlRepo *url.Repository, cache *cache.Cache) *Handler {
	return &Handler{
		urlRepo: urlRepo,
		cache: cache,
	}
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	// check redis
	longURL, err := h.cache.Get("url:" + code)
	if err != nil {
		u, err := h.urlRepo.FindByShortCode(code)
		if err != nil {
			http.Error(w, "url not found", http.StatusNotFound)
			return
		}

		longURL = u.LongURL

		h.cache.Set("url:"+code, longURL, time.Hour)

		h.urlRepo.RecordClick(u.ID)
	}

	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}

