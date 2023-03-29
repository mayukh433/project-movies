package models

import (
	"time"
)

type Movie struct {
	Isbn      string    `gorm:"primaryKey;not null;default:null" json:"isbn"`
	Title     string    `gorm:"not null;default:null" json:"title"`
	Director  string    `gorm:"not null;default:null" json:"director"`
	Genre     string    `gorm:"not null;default:null" json:"genre"`
	Rating    float32   `gorm:"not null;default:null" json:"rating"`
	CreatedAt time.Time `gorm:"not null;default:null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null;default:null" json:"updated_at"`
}
