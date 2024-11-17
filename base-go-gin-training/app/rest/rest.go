package rest

import (
	"base-gin/app/service"
	"base-gin/server"

	"github.com/gin-gonic/gin"
)

var (
	accountHandler   *AccountHandler
	personHandler    *PersonHandler
	publisherHandler *PublisherHandle
	authorHandler    *Author
	bookHandler      *BookHandler
	borrowHandler    *BorrowHandler
)

func SetupRestHandlers(app *gin.Engine) {
	handler := server.GetHandler()

	accountHandler = newAccountHandler(
		handler, service.GetAccountService(), service.GetPersonService())
	personHandler = newPersonHandler(
		handler, service.GetPersonService())
	publisherHandler = newPublisherHandler(
		handler, service.GetPublisherService())
	authorHandler = NewAuthorHandler(
		handler, service.GetAuthorService(),
	)
	bookHandler = NewBookHandler(
		handler, service.GetBookService(),
	)
	borrowHandler = NewBorrowHandler(
		handler, service.GetBorrowService(),
	)

	setupRoutes(app)
}

func setupRoutes(app *gin.Engine) {
	accountHandler.Route(app)
	personHandler.Route(app)
	publisherHandler.Route(app)
	authorHandler.Route(app)
	bookHandler.Route(app)
	borrowHandler.Route(app)
}
