package dao

import (
	"base-gin/app/domain"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Fullname  string             `gorm:"size:56;not null"`
	Gender    *domain.TypeGender `gorm:"type:enum('f','m');not null;" json:"gender"`
	BirthDate *time.Time        
}