package dto

import (
	"time"

	"github.com/bowodev/go-library-management/internal/model"
)

type Book struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	PublishDate time.Time `json:"publishDate" validate:"required"`
	Author      Author    `json:"author"`
}

type Books []Book

func (b Book) ToModel() model.Book {
	return model.Book{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		PublishDate: b.PublishDate,
		AuthorId:    b.Author.ID,
	}
}
