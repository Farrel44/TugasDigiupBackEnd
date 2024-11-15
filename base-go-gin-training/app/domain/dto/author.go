package dto

import (
	"base-gin/app/domain"
	"base-gin/app/domain/dao"
	"time"
)

type AuthorUpdateReq struct {
	ID           uint      `json:"-"`
	Fullname     string    `json:"fullname" binding:"required;max=56;min=4;not null"`
	Gender       string    `json:"gender" binding:"required;oneof=m f"` 
	BirthDateStr string    `json:"birth_date" binding:"required, datetime=2006-01-2"`
	BirthDate    time.Time `json:"-"`
}

func (v *AuthorUpdateReq) ToEntity() dao.Author {
	gender := domain.TypeGender(v.Gender)
	return dao.Author{
		Fullname:  v.Fullname,
		Gender:    &gender,
		BirthDate: &v.BirthDate,
	}
}

type AuthorCreateReq struct {
	Fullname     string    `json:"fullname" binding:"required;max=56;min=4;not null"`
	Gender       string    `json:"gender" binding:"required;oneof=m f"` 
	BirthDateStr string    `json:"birth_date" binding:"required, datetime=2006-01-2"`
	BirthDate    time.Time `json:"-"`                
}

func (v *AuthorCreateReq) ToEntity() dao.Author {
	gender := domain.TypeGender(v.Gender)
	return dao.Author{
		Fullname:  v.Fullname,
		Gender:    &gender,
		BirthDate: &v.BirthDate,
	}
}

type AuthorCreateResp struct {
	ID           uint      `json:"-"`
	Fullname     string    `json:"fullname" binding:"required;max=56;min=4;not null"`
	Gender       string    `json:"gender" binding:"required;oneof=m f"`
	BirthDate    time.Time `json:"-"`
	BirthDateStr string    `json:"birth_date" binding:"required, datetime=2006-01-2"`
}

func (o *AuthorCreateResp) FromEntity(item *dao.Author) {
	var gender string
	if item.Gender == nil {
		gender = "-"
	} else if *item.Gender == domain.GenderFemale {
		gender = "Female"
	} else {
		gender = "Male"
	}

	o.ID = uint(item.ID)
	o.Gender = gender
	o.Fullname = item.Fullname
	o.BirthDate = *item.BirthDate
	o.BirthDateStr = item.BirthDate.Format("2006-01-02")
}
