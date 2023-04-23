package controllers

import (
	"goPractice/errorhandler"
	"goPractice/models"
	"goPractice/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	CreateTaskCounter(ctx *gin.Context)
	GetCounterProgress(ctx *gin.Context)
	RemoveCounterTask(ctx *gin.Context)
}

type taskController struct {
	service      service.TaskService
	errorHandler errorhandler.ErrorHandler
}

func NewTaskController(service service.TaskService, errorhandler errorhandler.ErrorHandler) TaskController {
	return &taskController{
		service:      service,
		errorHandler: errorhandler,
	}
}

func (controller *taskController) CreateTaskCounter(ctx *gin.Context) {

	var request models.TaskCounterRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		controller.errorHandler.HandleErrorWithStatus(ctx, err, http.StatusBadRequest)
		return
	}

	taskId := controller.service.CreateTaskCounter(request)
	ctx.JSON(http.StatusOK, taskId)

}

func (controller *taskController) GetCounterProgress(ctx *gin.Context) {
	taskId, err := strconv.Atoi(ctx.Param("taskId"))
	if err != nil {
		controller.errorHandler.HandleError(err)
	}
	controller.service.GetCounterProgress(taskId)
}

func (controller *taskController) RemoveCounterTask(ctx *gin.Context) {
	taskId, err := strconv.Atoi(ctx.Param("taskId"))
	if err != nil {
		controller.errorHandler.HandleError(err)
	}
	controller.service.RemoveCounterTask(taskId)
}
