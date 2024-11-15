package integration_test

import (
	"base-gin/app/domain/dao"
	"base-gin/app/domain/dto"
	"base-gin/server"
	"base-gin/util"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Success(t *testing.T) {

	//pengujian
	req := dto.PublisherCreateReq{
		Name: util.RandomStringAlpha(6),
		City: util.RandomStringAlpha(10),
	}

	w := doTest("POST", server.RootPublisher, &req, createAuthAccessToken(dummyAdmin.Account.Username))

	//periksa hasil
	assert.Equal(t, 201, w.Code)

	var resp dto.SuccessResponse[dto.PublisherCreateResp]
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	if assert.Nil(t, err) {
		data := resp.Data
		assert.Greater(t, data.ID, 0)
		assert.Equal(t, req.Name, data.Name)
		assert.Equal(t, req.City, data.City)

		item, _ := publisherRepo.GetByID(uint(data.ID))
		if assert.NotNil(t, item) {
			assert.Equal(t, req.Name, item.Name)
			assert.Equal(t, req.City, item.City)
		}
	}
}
func TestPublisher_Delete_Success(t *testing.T) {
	o := dao.Publisher{
		Name: util.RandomStringAlpha(6),
		City: util.RandomStringAlpha(8),
	}
	_ = publisherRepo.Create(&o)

	w := doTest(
		"DELETE",
		fmt.Sprintf("%s/%d", server.RootPublisher, o.ID),
		nil,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)

	item, _ := publisherRepo.GetByID(o.ID)
	assert.Nil(t, item)
}

func Test_Update_Success(t *testing.T) {

}
