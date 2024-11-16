package dao

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"size:56;not null;"`
	Subtitle    string `gorm:"size:64:"`
	AuthorID    Author `gorm:"foreignKey:AuthorID;references:ID;"`
	PublisherID Publisher `gorm:"foreignKey:PublisherID;references:ID;"`
}
