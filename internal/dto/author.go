package dto

import "time"

type Author struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	BirthDate time.Time `json:"birth_date"`
}

type Authors []Author
