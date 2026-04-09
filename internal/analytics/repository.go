package analytics

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return  &Repository{db: db}
}

func (r *Repository) GetClickCount(urlID int) (int, error) {
	var count int
	err := r.db.Get(&count, `SELECT COUNT(*) FROM clicks WHERE url_id = $1`, urlID)
	return  count, err
}

