package dto

import "base-gin/app/domain/dao"

type BookCreateReq struct {
	Title       string `json:"title" binding:"size:56;not null;"`
	Subtitle    string `json:"subtitle" binding:"size:64:"`
	AuthorID    uint   `json:"AuthorID" binding:"required"`
	PublisherID uint   `json:"PublisherID" binding:"required"`
}

type BookCreateResp struct {
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle"`
	AuthorID    uint   `json:"AuthorID"`
	PublisherID uint   `json:"PublisherID"`
}

type BookUpdateReq struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

func (o *BookCreateReq) FromEntity(item *dao.Book) {
	o.Title = item.Title
	o.Subtitle = item.Subtitle
}

func (o *BookCreateResp) BookRes(item *dao.Book) {
	o.Title = item.Title
	o.Subtitle = item.Subtitle
}

func (o *BookCreateReq) ToEntity() dao.Book {
	var item dao.Book
	o.Title = item.Title
	o.Subtitle = item.Subtitle
	return item
}

func (o *BookUpdateReq) UpdateBook(item *dao.Book, id uint) {
	o.Title = item.Title
	o.Subtitle = item.Subtitle
}
