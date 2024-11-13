package rest

import (
	"base-gin/app/domain/dto"
	"base-gin/app/service"
	"base-gin/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublisherHandle struct {
	hr      *server.Handler
	service *service.PublisherService
}

func newPublisherHandler(hr *server.Handler, publisherService *service.PublisherService) *PublisherHandle {
	return &PublisherHandle{hr: hr, service: publisherService}
}

func (h *PublisherHandle) Route(app *gin.Engine) {
	grp := app.Group("/v1/publishers")
	grp.POST("", h.hr.AuthAccess(), h.Create)
}

func (h *PublisherHandle) Create(c *gin.Context) {
	var req dto.PublisherCreateReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.hr.BindingError(err)
		return
	}

	data, err := h.service.Create(&req)
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}

	c.JSON(http.StatusCreated, dto.SuccessResponse[*dto.PublisherCreateResp]{
		Success: true,
		Message: "Data penerbit berhasil disimpan",
		Data:    data,
	})
}
