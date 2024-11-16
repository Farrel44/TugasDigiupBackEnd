package rest

import (
	"base-gin/app/domain/dto"
	"base-gin/app/service"
	"base-gin/exception"
	"base-gin/server"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Author struct {
	hr      *server.Handler
	service *service.AuthorService
}

func NewAuthorHandler(handler *server.Handler, authorService *service.AuthorService) *Author {
	return &Author{hr: handler, service: authorService}
}

func (h *Author) Route(app *gin.Engine) {
	grp := app.Group(server.RootAuthor)
	app.RedirectTrailingSlash = false
	grp.GET("", h.getList)
	grp.GET("/:id", h.getByID)
	grp.POST("", h.hr.AuthAccess(), h.create)
	grp.DELETE("/:id", h.hr.AuthAccess(), h.delete)
	grp.PUT("/:id", h.hr.AuthAccess(), h.update)
}

func (h *Author) create(c *gin.Context) {
	var req dto.AuthorCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.hr.BindingError(err)
		return
	}
	resp, err := h.service.Create(&req)
	if err != nil {
		h.hr.ErrorInternalServer(c, err)
		return
	}
	c.JSON(http.StatusAccepted, dto.SuccessResponse[any]{
		Success: true,
		Message: "data berhasil disimpan",
		Data:    resp,
	})
}

func (h *Author) getList(c *gin.Context) {
	var req dto.Filter
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(h.hr.BindingError(err))
		return
	}
	data, err := h.service.GetList(&req)
	if err != nil {
		switch {
		case errors.Is(err, exception.ErrDataNotFound):
			c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		default:
			h.hr.ErrorInternalServer(c, err)
		}
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[[]dto.AuthorCreateResp]{
		Success: true,
		Message: "Daftar Penerbit",
		Data:    data,
	})
}

func (h *Author) getByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	author, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, h.hr.ErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse[dto.AuthorCreateResp]{
		Success: true,
		Message: "Author details",
		Data:    author,
	})
}

func (h *Author) update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	var req dto.AuthorUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("JSON binding/validation error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data", "details": err.Error()})
		return
	}
	err = h.service.Update(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update author"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Author updated successfully"})
}

func (h *Author) delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete author"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Author deleted successfully"})
}
