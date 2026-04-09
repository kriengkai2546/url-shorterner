package url

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(userID int, longURL, shortCode string) (*URL, error) {
	u := &URL{}
	query := `
		INSERT INTO urls (user_id, long_url, short_code)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, long_url, short_code, created_at
	`

	err := r.db.QueryRowx(query, userID, longURL, shortCode).StructScan(u)
	return  u, err
}

func (r *Repository) FindByShortCode(code string) (*URL, error) {
	u := &URL{}
	query := `SELECT id, user_id, long_url, short_code, created_at FROM urls WHERE short_code = $1`
	err := r.db.Get(u, query, code)
	return u, err
}

func (r *Repository) FindByUserID(userID int) ([]URL, error) {
	var urls []URL
	query := `SELECT id, user_id, long_url, short_code, created_at FROM urls WHERE user_id = $1 ORDER BY created_at DESC`
	err := r.db.Select(&urls, query, userID)
	return urls, err
}

func (r *Repository) Delete(id, userID int) error {
	query := `DELETE FROM urls WHERE id = $1 AND user_id = $2`
	_, err := r.db.Exec(query, id, userID)
	return  err
}

func (r *Repository) RecordClick(urlID int) error {
	_, err := r.db.Exec(`INSERT INTO clicks (url_id) VALUES ($1)`, urlID)
	return err
}