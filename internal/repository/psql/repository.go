package psql

import (
	"github.com/AndrewMislyuk/CRUD-languages/internal/domain"
	"database/sql"
)

type Language interface {
	Create(language domain.Language) (string, error)
	GetByID(id string) (domain.Language, error)
	GetAll() ([]domain.Language, error)
	Delete(id string) error
	Update(id string, inp domain.UpdateLanguageInput) error
}

type User interface {
	CreateUser(user domain.SignUpInput) (string, error)
	GetUser(email, password string) (string, error)
}

type Repository struct {
	Language
	User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Language: NewLanguagesRepository(db),
		User: NewAuthRepository(db),
	}
}
