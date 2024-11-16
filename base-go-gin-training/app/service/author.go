package service

import (
	"base-gin/app/domain"
	"base-gin/app/domain/dto"
	"base-gin/app/repository"
	"base-gin/exception"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type AuthorService struct {
	repo *repository.AuthorRepository
}

func newAuthorService(authorRepo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: authorRepo}
}

func (s *AuthorService) Create(params *dto.AuthorCreateReq) (*dto.AuthorCreateResp, error) {
	newItem := params.ToEntity()

	err := s.repo.Create(&newItem)
	if err != nil {
		return nil, err
	}

	var resp dto.AuthorCreateResp
	resp.FromEntity(&newItem)

	return &resp, nil
}

func (s *AuthorService) GetList(params *dto.Filter) ([]dto.AuthorCreateResp, error) {
	var resp []dto.AuthorCreateResp

	items, err := s.repo.GetList(params)
	if err != nil {
		return nil, err
	}
	if len(items) < 1 {
		return nil, exception.ErrUserNotFound
	}

	for _, item := range items {
		var t dto.AuthorCreateResp
		t.FromEntity(&item)

		resp = append(resp, t)
	}

	return resp, nil
}

func (s *AuthorService) GetByID(id uint) (dto.AuthorCreateResp, error) {
	var resp dto.AuthorCreateResp

	item, err := s.repo.GetByID(id)
	if err != nil {
		return resp, err
	}
	if item == nil {
		return resp, exception.ErrUserNotFound
	}
	resp.FromEntity(item)

	return resp, nil
}

func (s *AuthorService) Update(id uint, req *dto.AuthorUpdateReq) error {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		return err
	}

	author, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	author.Fullname = req.Fullname

	if req.Gender != "" {
		gender := domain.TypeGender(req.Gender)
		author.Gender = &gender
	}

	err = s.repo.Update(author)
	if err != nil {
		return err
	}

	return nil
}
func (s *AuthorService) Delete(id uint) error {
	if id <= 0 {
		return exception.ErrDataNotFound
	}
	return s.repo.Delete(id)
}
