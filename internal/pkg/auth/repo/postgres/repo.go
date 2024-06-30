package postgres

import "database/sql"

type AuthRepositoryImpl struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{DB: db}
}

func (r *AuthRepositoryImpl) CheckUser(name, password string) (int, error) {
	var userID int
	query := `SELECT id FROM "user" WHERE name = $1 AND password_hash = $2`
	err := r.DB.QueryRow(query, name, password).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (r *AuthRepositoryImpl) CreateUser(name, password string) (int, error) {
	var userID int
	query := `INSERT INTO "user" (name, password_hash) VALUES ($1, $2) RETURNING id`
	err := r.DB.QueryRow(query, name, password).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
