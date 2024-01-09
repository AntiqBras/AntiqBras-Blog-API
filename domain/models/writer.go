package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Writer struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:400;not null;" json:"password"`
	Author string `gorm:"size:150;not null;" json:"author"`

	Posts []Post `json:"posts"`
	
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (w *Writer) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(w.Password), 10)
	if err != nil {
		return err
	}

	w.Password = string(hashedPassword)
	w.Username = strings.ReplaceAll(w.Username, " ", "")
	return nil
}

func (w *Writer) PrepareGive() {
	w.Password = ""
}
