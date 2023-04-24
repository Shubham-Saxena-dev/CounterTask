package controllers

import (
	"goPractice/errorhandler"
	"goPractice/models"
	"goPractice/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	CreateTaskCounter(ctx *gin.Context)
	GetTaskProgress(ctx *gin.Context)
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

func (controller *taskController) GetTaskProgress(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	progress := controller.service.GetTaskProgress(taskId)
	ctx.JSON(http.StatusOK, progress)

}

func (controller *taskController) RemoveCounterTask(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	status := controller.service.RemoveCounterTask(taskId)
	ctx.JSON(http.StatusOK, status)
}
