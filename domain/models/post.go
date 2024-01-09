package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID uint `gorm:"primaryKey"`
	HeroImage string `gorm:"not null" json:"hero_image"`
	Title string `gorm:"not null;unique" json:"title"`
	Subtitle string `gorm:"not null" json:"subtitle"`
	Slug string `gorm:"not null;unique" json:"slug"`
	Content string `gorm:"not null" json:"content"`

	WriterName string `json:"writer_name"`
	WriterID uuid.UUID `json:"writer_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
