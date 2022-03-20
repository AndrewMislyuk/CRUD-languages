package domain

import "errors"

var ErrLanguageNotFound = errors.New("language not found")

type Language struct {
	Id             string `json:"id"`
	Title          string `json:"title"`
	Rating         int64  `json:"rating"`
	Developer      string `json:"developer"`
	DateOfCreation int32  `json:"date_of_creation"`
}

type UpdateLanguageInput struct {
	Title          *string `json:"title"`
	Rating         *int64  `json:"rating"`
	Developer      *string `json:"developer"`
	DateOfCreation *int32  `json:"date_of_creation"`
}
