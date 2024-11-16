package dto

import (
	"base-gin/app/domain"
	"base-gin/app/domain/dao"
	"fmt"
	"time"
)

type AuthorUpdateReq struct {
	ID           uint      `json:"-"`
	Fullname     string    `json:"fullname" binding:"required,max=255"`
	Gender       string    `json:"gender" binding:"omitempty,oneof=m f"`
	BirthDateStr string    `json:"birth_date" binding:"required,datetime=2006-01-02"`
	BirthDate    time.Time `json:"-"`
	Age          int       `json:"age"`
}

func (v *AuthorUpdateReq) ToEntity() (dao.Author, error) {
	birthDate, err := time.Parse("2006-01-02", v.BirthDateStr)
	if err != nil {
		return dao.Author{}, fmt.Errorf("invalid date format, expected YYYY-MM-DD: %v", err)
	}

	gender := domain.TypeGender(v.Gender)

	return dao.Author{
		Fullname:  v.Fullname,
		Gender:    &gender,
		BirthDate: &birthDate,
	}, nil
}

type AuthorCreateReq struct {
	Fullname     string    `json:"fullname" binding:"required,min=4,max=56"`
	Gender       string    `json:"gender" binding:"required,oneof=m f"`
	BirthDateStr string    `json:"birth_date" binding:"required,datetime=2006-01-02"`
	BirthDate    time.Time `json:"-"`
	Age          int       `json:"age"`
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
	Fullname     string    `json:"fullname" binding:"required;max=56;min=4"`
	Gender       string    `json:"gender" binding:"required;oneof=m f"`
	BirthDate    time.Time `json:"-"`
	BirthDateStr string    `json:"birth_date" binding:"required;datetime=2006-01-2"`
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
