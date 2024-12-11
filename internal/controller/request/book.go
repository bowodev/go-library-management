package request

import (
	"time"

	"github.com/bowodev/go-library-management/internal/dto"
)

// CreateBook
type CreateBook struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishDate string `json:"publishDate"`
	AuthorId    int64  `json:"authorId"`
}

func (r CreateBook) ToDTO() (dto.Book, error) {
	publishDate, err := time.Parse("2006-01-02", r.PublishDate)
	if err != nil {
		return dto.Book{}, err
	}

	return dto.Book{
		Title:       r.Title,
		Description: r.Description,
		Author: dto.Author{
			ID: r.AuthorId,
		},
		PublishDate: publishDate,
	}, nil
}

// UpdateBook

type UpdateBook struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PublishDate string `json:"publishDate"`
	AuthorId    int64  `json:"authorId"`
}

func (r UpdateBook) ToDTO() (dto.Book, error) {
	publishDate, err := time.Parse("2006-01-02", r.PublishDate)
	if err != nil {
		return dto.Book{}, err
	}

	return dto.Book{
		Title:       r.Title,
		Description: r.Description,
		Author: dto.Author{
			ID: r.AuthorId,
		},
		PublishDate: publishDate,
	}, nil
}
