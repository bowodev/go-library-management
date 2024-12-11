package request

import (
	"time"

	"github.com/bowodev/go-library-management/internal/dto"
)

// CreateAuthor
type CreateAuthor struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	BirthDate string `json:"birthDate"`
}

func (r CreateAuthor) ToDTO() (dto.Author, error) {
	birthDate, err := time.Parse("2006-01-02", r.BirthDate)
	if err != nil {
		return dto.Author{}, err
	}

	return dto.Author{
		ID:        r.ID,
		Name:      r.Name,
		Bio:       r.Bio,
		BirthDate: birthDate,
	}, nil
}

// UpdateAuthor
type UpdateAuthor struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	BirthDate string `json:"birthDate"`
}

func (r UpdateAuthor) ToDTO() (dto.Author, error) {
	birthDate, err := time.Parse("2006-01-02", r.BirthDate)
	if err != nil {
		return dto.Author{}, err
	}

	return dto.Author{
		ID:        r.ID,
		Name:      r.Name,
		Bio:       r.Bio,
		BirthDate: birthDate,
	}, nil
}
