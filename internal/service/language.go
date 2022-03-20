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
