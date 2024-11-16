package dao

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string 
	Subtitle string 
}