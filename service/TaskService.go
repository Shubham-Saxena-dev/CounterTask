package service

import (
	"goPractice/errorhandler"
	"goPractice/models"
	"goPractice/repository"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTaskCounter(request models.TaskCounterRequest) uuid.UUID
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

func (service *taskService) CreateTaskCounter(request models.TaskCounterRequest) uuid.UUID {
	request.TaskId = uuid.New()
	return service.database.SaveTaskCounters(request)

}

func (service *taskService) GetCounterProgress(taskId int) {

}

func (service *taskService) RemoveCounterTask(taskId int) {

}
