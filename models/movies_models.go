package models

import (
	"time"
	//"go.mongodb.org/mongo-driver/bson/primitive"
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

//type UpdateMovie struct {
//	Isbn      string    `json:"title"`
//	Title     string    `json:"title" bson:"content"`
//	Director  Director  `json:"director" bson:"director"`
//	Genre     string    `json:"genre" bson:"genre"`
//	Rating    float32   `json:"rating" bson:"rating"`
//	CreatedAt time.Time `json:"created_at" bson:"created_at"`
//	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
//}
