package integration_test

import (
	"base-gin/app/domain/dto"
	"base-gin/server"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBorrow_Create_Success(t *testing.T) {
	borrowDate, _ := time.Parse("2006-01-02", "2024-11-01")
	returnDate, _ := time.Parse("2006-01-02", "2024-11-15")
	req := dto.BorrowBookReq{
		BookId:      1,
		PublisherID: 1,
		BorrowDate:  &borrowDate,
		ReturnDate:  &returnDate,
	}

	w := doTest("POST", server.RootBorrow, req, createAuthAccessToken(dummyAdmin.Account.Username))

	assert.Equal(t, 202, w.Code)
}
