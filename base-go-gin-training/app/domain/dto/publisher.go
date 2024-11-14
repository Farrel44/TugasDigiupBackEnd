package dto

import "base-gin/app/domain/dao"

type PublisherCreateReq struct {
    Name string `json:"name" binding:"required,max=48,min=6"`
    City string `json:"city" binding:"required,max=32,min=2"`
}

func (v *PublisherCreateReq) ToEntity() dao.Publisher {
	return dao.Publisher{
		Name : v.Name,
		City : v.City,
	}
}

type PublisherCreateResp struct {
	ID int `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func (v *PublisherCreateResp) FromEntity(item *dao.Publisher){
	v.ID = int(item.ID)
	v.Name = item.Name
	v.City = item.City

}