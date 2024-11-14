package integration_test

import (
	"base-gin/app/domain/dto"
	"base-gin/util"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Success(t *testing.T) {
	req := dto.PublisherCreateReq{
		Name: util.RandomStringAlpha(6),
		City: util.RandomStringAlpha(10),
	}

	w := doTest("POST", "/v1/publishers", &req,
		createAuthAccessToken(dummyAdmin.Account.Username))
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
