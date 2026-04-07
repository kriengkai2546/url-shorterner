package url

import "time"

type URL struct {
	ID        int       `db:"id" json:"id"`
	UserID    int       `db:"user_id" json:"user_id"`
	LongURL   string    `db:"long_url" json:"long_url"`
	ShortCode string    `db:"short_code" json:"short_code"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type CreateURLRequest struct {
	LongURL string `json:"long_url"`
}

type CreateURLResponse struct {
	ShortCode string `json:"short_code"`
	ShortURL string `json:"short_url"`
	LongURL string `json:"long_url"`
}