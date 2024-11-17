package integration_test

import (
	"base-gin/app/domain/dto"
	"base-gin/server"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBook_Create_Success(t *testing.T) {
	req := dto.BookCreateReq{
		Title:    "Sample Book",
		Subtitle: "Lorem12",
		AuthorID: 1,
	}
	w := doTest("POST", server.RootBook+"/", req, createAuthAccessToken(dummyAdmin.Account.Username))
	assert.Equal(t, 200, w.Code)
}
