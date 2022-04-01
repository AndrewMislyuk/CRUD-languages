package service

import (
	"github.com/AndrewMislyuk/CRUD-languages/internal/domain"
)

type LanguageRepository interface {
	Create(language domain.Language) (string, error)
	GetByID(id string) (domain.Language, error)
	GetAll() ([]domain.Language, error)
	Delete(id string) error
	Update(id string, inp domain.UpdateLanguageInput) error
}

type Languages struct {
	repo LanguageRepository
}

func NewLanguages(repo LanguageRepository) *Languages {
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
