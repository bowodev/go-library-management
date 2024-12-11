package dto

import (
	"time"

	"github.com/bowodev/go-library-management/internal/model"
)

type Author struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	BirthDate time.Time `json:"birthDate"`
}

type Authors []Author

func (a Author) ToModel() model.Author {
	return model.Author{
		ID:        a.ID,
		Name:      a.Name,
		Bio:       a.Bio,
		BirthDate: a.BirthDate,
	}
}

func (a Authors) ToModels() []model.Author {
	authors := []model.Author{}
	for _, v := range a {
		authors = append(authors, v.ToModel())
	}

	return authors
}

func (a *Author) FromModel(m model.Author) {
	if a == nil {
		a = &Author{}
	}

	a.ID = m.ID
	a.Name = m.Name
	a.Bio = m.Bio
	a.BirthDate = m.BirthDate
}

func (a *Authors) FromModels(m []model.Author) {
	authors := []Author{}

	for _, v := range m {
		authors = append(authors, Author{
			ID:        v.ID,
			Name:      v.Name,
			Bio:       v.Bio,
			BirthDate: v.BirthDate,
		})
	}

	*a = authors
}
