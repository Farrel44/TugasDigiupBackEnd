package repository

import "gorm.io/gorm"

type BookRepository struct {
	db *gorm.DB
}

func newBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}
