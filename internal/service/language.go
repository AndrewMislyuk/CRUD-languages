package service

import (
	"github.com/AndrewMislyuk/CRUD-languages/internal/domain"
	repo "github.com/AndrewMislyuk/CRUD-languages/internal/repository/psql"
)

type Languages struct {
	repo repo.Language
}

func NewLanguagesService(repo repo.Language) *Languages {
	return &Languages{
		repo: repo,
	}
}

func (l *Languages) GetByID(id string) (domain.Language, error) {
	return l.repo.GetByID(id)
}

func (l *Languages) Update(id string, inp domain.UpdateLanguageInput) error {
	return l.repo.Update(id, inp)
}

func (l *Languages) Delete(id string) error {
	return l.repo.Delete(id)
}

func (l *Languages) Create(language domain.Language) (string, error) {
	return l.repo.Create(language)
}

func (l *Languages) GetAll() ([]domain.Language, error) {
	return l.repo.GetAll()
}
