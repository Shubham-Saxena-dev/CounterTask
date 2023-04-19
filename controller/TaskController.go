package controller

import (
	"goPractice/service"
	"goPractice/errorHandler"

)

type Controller interface {
	createTaskCounter()
	getCounterProgress()
	removeCounterTask()
}

type controller struct {
	service      service.TaskService
	errorHandler errorhandler.ErrorHandler
}

func NewController(service service.TaskService, errorhandler errorhandler.ErrorHandler) Controller {
	return &controller{
		service:      service,
		errorHandler: errorhandler,
	}
}

func (controller *controller) createTaskCounter() {

}

func (controller *controller) getCounterProgress() {

}

func (controller *controller) removeCounterTask() {

}
