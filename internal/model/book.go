package model

import (
	"time"
)

type Book struct {
	ID          int64     `gorm:"id"`
	Title       string    `gorm:"title"`
	Description string    `gorm:"description"`
	PublishDate time.Time `gorm:"column:publish_date"`
	AuthorId    int64     `gorm:"column:author_id"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
