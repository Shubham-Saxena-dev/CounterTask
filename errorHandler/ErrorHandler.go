package errorhandler

import "github.com/gin-gonic/gin"

type ErrorHandler interface {
	HandleErrorWithStatus(ctx *gin.Context, err error, status int)
	HandleError(err error)
}

type errorHandler struct {
}

func NewErrorHandler() ErrorHandler {
	return &errorHandler{}
}

func (errorHandler *errorHandler) HandleErrorWithStatus(ctx *gin.Context, err error, status int) {

}

func (errorHandler *errorHandler) HandleError(err error) {

}
