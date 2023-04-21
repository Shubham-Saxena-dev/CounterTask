package service

import (
	"goPractice/errorhandler"
	"goPractice/repository"
)

type TaskService interface {
	CreateTaskCounter(start, end int)
	GetCounterProgress(taskId int)
	RemoveCounterTask(int)
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

func (service *taskService) CreateTaskCounter(start, end int) {

}

func (service *taskService) GetCounterProgress(taskId int) {

}

func (service *taskService) RemoveCounterTask(taskId int) {

}
