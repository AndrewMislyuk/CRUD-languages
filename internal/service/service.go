package service

import (
	"github.com/AndrewMislyuk/CRUD-languages/internal/domain"
	repo "github.com/AndrewMislyuk/CRUD-languages/internal/repository/psql"
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
	GenerateToken(email, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Service struct {
	Language
	User
}

func NewService(repo *repo.Repository) *Service {
	return &Service{
		Language: NewLanguagesService(repo.Language),
		User:     NewAuthService(repo.User),
	}
}
