package repository

import (
	"base-gin/app/domain/dao"
	"base-gin/exception"
	"base-gin/storage"
	"errors"

	"gorm.io/gorm"
)

type PublisherRespository struct {
	db *gorm.DB
}

func newPublisherRepo(db *gorm.DB) *PublisherRespository {
	return &PublisherRespository{db: db}
}

func (r *PublisherRespository) Create(newItem *dao.Publisher) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	tx := r.db.WithContext(ctx).Create(&newItem)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *PublisherRespository) GetByID(id uint) (*dao.Publisher, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	var item dao.Publisher
	tx := r.db.WithContext(ctx).First(&item, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, exception.ErrUserNotFound
		}

		return nil, tx.Error
	}

	return &item, nil
}
