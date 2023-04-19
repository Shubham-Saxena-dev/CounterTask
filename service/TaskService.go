package service

import (
	errorhandler "goPractice/errorHandler"
	"goPractice/repository"
)

type TaskService interface {
	createTaskCounter()
	getCounterProgress()
	removeCounterTask()
}

type taskService struct {
	database     repository.TaskRepository
	errorHandler errorhandler.ErrorHandler
}

func NewTaskService(database repository.TaskRepository, errorHandler errorhandler.ErrorHandler) TaskService {
	return &taskService{
		database:     database,
		errorHandler: errorHandler,
	}
}

func (service *taskService) createTaskCounter() {

}

func (service *taskService) getCounterProgress() {

}

func (service *taskService) removeCounterTask() {

}

