package errorhandler

type ErrorHandler interface {
	handleError(err error)
}

type errorHandler struct {
}

func NewErrorHandler() ErrorHandler {
	return &errorHandler{}
}

func (errorHandler *errorHandler) handleError(err error) {

}
