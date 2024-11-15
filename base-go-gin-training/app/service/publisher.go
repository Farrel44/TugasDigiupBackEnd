package service

import (
	"base-gin/app/domain/dto"
	"base-gin/app/repository"
	"base-gin/exception"
)

type PublisherService struct {
	repo *repository.PublisherRespository
}

// Endpoint

func newPublisherService(publisherRepo *repository.PublisherRespository) *PublisherService {
	return &PublisherService{repo: publisherRepo}
}

func (s *PublisherService) Create(params *dto.PublisherCreateReq) (*dto.PublisherCreateResp, error) {
	newItem := params.ToEntity()

	err := s.repo.Create(&newItem)
	if err != nil {
		return nil, err
	}

	var resp dto.PublisherCreateResp
	resp.FromEntity(&newItem)

	return &resp, nil
}

func (s *PublisherService) GetList(params *dto.Filter) ([]dto.PublisherCreateResp, error) {
	var resp []dto.PublisherCreateResp

	items, err := s.repo.GetList(params)
	if err != nil {
		return nil, err
	}
	if len(items) < 1 {
		return nil, exception.ErrUserNotFound
	}

	for _, item := range items {
		var t dto.PublisherCreateResp
		t.FromEntity(&item)

		resp = append(resp, t)
	}

	return resp, nil
}

func (s *PublisherService) GetByID(id uint) (dto.PublisherCreateResp, error) {
	var resp dto.PublisherCreateResp

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

func (s *PublisherService) Update(params *dto.PublisherUpdateReq) error {
	if params.ID <= 0 {
		return exception.ErrDataNotFound
	}

	return s.repo.Update(params)
}

func (s *PublisherService) Delete(id uint) error {
	if id <= 0 {
		return exception.ErrDataNotFound
	}
	return s.repo.Delete(id)
}
