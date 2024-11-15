package dao

import (
	"base-gin/app/domain"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID int ``
	Fullname  string             `gorm:"size:56;not null;"`
	Gender    *domain.TypeGender `gorm:"type:enum('f','m');not null;"`
	BirthDate *time.Time
}
