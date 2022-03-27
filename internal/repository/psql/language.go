package psql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/AndrewMislyuk/CRUD-languages/internal/domain"
)

type Languages struct {
	db *sql.DB
}

func NewLanguages(db *sql.DB) *Languages {
	return &Languages{
		db: db,
	}
}

func (l *Languages) GetByID(ctx context.Context, id string) (domain.Language, error) {
	rows, err := l.db.Query("SELECT * FROM languages WHERE id = $1", id)
	if err != nil {
		return domain.Language{}, err
	}

	var language domain.Language
	for rows.Next() {
		if err := rows.Scan(&language.Id, &language.Title, &language.Rating, &language.Developer, &language.DateOfCreation); err != nil {
			return domain.Language{}, err
		}
	}

	return language, rows.Err()
}

func (l *Languages) Update(ctx context.Context, id string, inp domain.UpdateLanguageInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if inp.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *inp.Title)
		argId++
	}

	if inp.Rating != nil {
		setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
		args = append(args, *inp.Rating)
		argId++
	}

	if inp.Developer != nil {
		setValues = append(setValues, fmt.Sprintf("developer=$%d", argId))
		args = append(args, *inp.Developer)
		argId++
	}

	if inp.DateOfCreation != nil {
		setValues = append(setValues, fmt.Sprintf("date_of_creation=$%d", argId))
		args = append(args, *inp.DateOfCreation)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE languages SET %s WHERE id=$%d", setQuery, argId)
	args = append(args, id)

	_, err := l.db.Exec(query, args...)
	return err
}

func (l *Languages) Delete(ctx context.Context, id string) error {
	_, err := l.db.Exec("DELETE FROM languages WHERE id = $1", id)

	return err
}

func (l *Languages) Create(ctx context.Context, language domain.Language) error {
	_, err := l.db.Exec("INSERT INTO languages(title, rating, developer, date_of_creation) values ($1, $2, $3, $4)",
		language.Title, language.Rating, language.Developer, language.DateOfCreation)

	return err
}

func (l *Languages) GetAll(ctx context.Context) ([]domain.Language, error) {
	rows, err := l.db.Query("SELECT * FROM languages")
	if err != nil {
		return nil, err
	}

	languages := make([]domain.Language, 0)
	for rows.Next() {
		var language domain.Language
		if err := rows.Scan(&language.Id, &language.Title, &language.Rating, &language.Developer, &language.DateOfCreation); err != nil {
			return nil, err
		}

		languages = append(languages, language)
	}

	return languages, rows.Err()
}
