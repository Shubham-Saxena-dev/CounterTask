package service

import (
	"goPractice/errorhandler"
	"goPractice/models"
	"goPractice/repository"

	"github.com/google/uuid"
)

type TaskService interface {
	CreateTaskCounter(request models.TaskCounterRequest) string
	GetTaskProgress(taskId string) int
	RemoveCounterTask(string) string
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

func (service *taskService) CreateTaskCounter(request models.TaskCounterRequest) string {
	request.TaskId = uuid.NewString()
	return service.database.SaveTaskCounters(request)

}

func (service *taskService) GetTaskProgress(taskId string) int {
	return service.database.GetTaskProgress(taskId)
}

func (service *taskService) RemoveCounterTask(taskId string) string {
	return service.database.RemoveCounterTask(taskId)

}
