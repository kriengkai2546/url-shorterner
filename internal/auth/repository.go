package auth

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(email, hashedPassword string) (*User, error) {
	user := &User{}
	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id, email, password, created_at
	`
	err := r.db.QueryRowx(query, email, hashedPassword).StructScan(user)
	if err != nil {
		return nil, err
	}
	return  user, nil
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	user := &User{}
	query := `SELECT id, email, password, created_at FROM users WHERE email = $1`
	err := r.db.Get(user, query, email) 
	if err != nil {
		return nil, err
	}
	return user, nil
}