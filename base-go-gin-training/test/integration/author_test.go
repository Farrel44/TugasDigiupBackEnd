package integration_test

import (
	"base-gin/app/domain/dto"
	"base-gin/server"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthor_Create_Success(t *testing.T) {
	req := dto.AuthorCreateReq{
		Fullname:     "John Doe",
		Gender:       "m",
		BirthDateStr: "2000-01-01",
	}

	w := doTest("POST", server.RootAuthor, req, createAuthAccessToken(dummyAdmin.Account.Username))
	assert.Equal(t, 202, w.Code)
}

func TestAuthor_GetList_Success(t *testing.T) {
	accessToken := createAuthAccessToken(dummyAdmin.Account.Username)
	w := doTest("GET", server.RootAuthor, nil, accessToken)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}

func TestAuthor_GetByID_Success(t *testing.T) {
	accessToken := createAuthAccessToken(dummyAdmin.Account.Username)
	w := doTest("GET", server.RootAuthor+"/1", nil, accessToken)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}
func TestAuthor_Update_Success(t *testing.T) {
	req := dto.AuthorUpdateReq{
		Fullname:     "John Doe Updated",
		Gender:       "m",
		BirthDateStr: "1990-01-01",
	}
	accessToken := createAuthAccessToken(dummyAdmin.Account.Username)
	w := doTest("PUT", server.RootAuthor+"/1", req, accessToken)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Author updated successfully")
}

func TestAuthor_Delete_Success(t *testing.T) {
	accessToken := createAuthAccessToken(dummyAdmin.Account.Username)
	w := doTest("DELETE", server.RootAuthor+"/1", nil, accessToken)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Author deleted successfully")
}
