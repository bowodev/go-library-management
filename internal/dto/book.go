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

func (b Books) ToModels() []model.Book {
	books := []model.Book{}
	for _, v := range b {
		books = append(books, v.ToModel())
	}

	return books
}

func (b *Book) FromModel(m model.Book) {
	if b == nil {
		b = &Book{}
	}

	b.ID = m.ID
	b.Title = m.Title
	b.Description = m.Description
	b.PublishDate = m.PublishDate
	b.Author = Author{
		ID:        m.AuthorId,
		Name:      m.Author.Name,
		Bio:       m.Author.Bio,
		BirthDate: m.Author.BirthDate,
	}
}

func (b *Books) FromModels(m []model.Book) {
	books := Books{}

	for _, v := range m {
		books = append(books, Book{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			PublishDate: v.PublishDate,
			Author: Author{
				ID:        v.AuthorId,
				Name:      v.Author.Name,
				Bio:       v.Author.Bio,
				BirthDate: v.Author.BirthDate,
			},
		})
	}

	*b = books
}
