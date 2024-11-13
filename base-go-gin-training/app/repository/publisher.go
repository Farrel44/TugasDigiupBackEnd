package repository

import (
	"base-gin/app/domain/dao"
	"base-gin/storage"

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
