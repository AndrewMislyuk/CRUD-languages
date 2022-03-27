package service

import (
	"context"

	"github.com/AndrewMislyuk/CRUD-languages/internal/domain"
)

type LanguageRepository interface {
	Create(ctx context.Context, language domain.Language) error
	GetByID(ctx context.Context, id string) (domain.Language, error)
	GetAll(ctx context.Context) ([]domain.Language, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, inp domain.UpdateLanguageInput) error
}

type Languages struct {
	repo LanguageRepository
}

func NewLanguages(repo LanguageRepository) *Languages {
	return &Languages{
		repo: repo,
	}
}

func (l *Languages) GetByID(ctx context.Context, id string) (domain.Language, error) {
	return l.repo.GetByID(ctx, id)
}

func (l *Languages) Update(ctx context.Context, id string, inp domain.UpdateLanguageInput) error {
	return l.repo.Update(ctx, id, inp)
}

func (l *Languages) Delete(ctx context.Context, id string) error {
	return l.repo.Delete(ctx, id)
}

func (l *Languages) Create(ctx context.Context, language domain.Language) error {
	return l.repo.Create(ctx, language)
}

func (l *Languages) GetAll(ctx context.Context) ([]domain.Language, error) {
	return l.repo.GetAll(ctx)
}
