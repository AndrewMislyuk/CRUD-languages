package psql

import (
	"database/sql"

	"github.com/AndrewMislyuk/CRUD-languages/internal/domain"
)

type Auth struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *Auth {
	return &Auth{db: db}
}

func (a *Auth) CreateUser(user domain.SignUpInput) (string, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return "", err
	}

	var id string
	row, err := tx.Prepare("INSERT INTO users(name, email, password_hash) values($1, $2, $3) RETURNING id")
	if err != nil {
		return "", err
	}

	defer row.Close()

	if err = row.QueryRow(user.Name, user.Email, user.Password).Scan(&id); err != nil {
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return id, nil
}

func (a *Auth) GetUser(email, password string) (string, error) {
	var userId string
	rows, err := a.db.Query("SELECT id FROM users WHERE email = $1 AND password_hash = $2", email, password)
	if err != nil {
		return "", err
	}

	for rows.Next() {
		if err := rows.Scan(&userId); err != nil {
			return "", err
		}
	}

	return userId, rows.Err()
}
