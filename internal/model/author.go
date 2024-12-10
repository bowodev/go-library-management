package model

import "time"

type Author struct {
	ID        int64     `gorm:"id"`
	Name      string    `gorm:"name"`
	Bio       string    `gorm:"bio"`
	BirthDate time.Time `gorm:"column:birth_date"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
