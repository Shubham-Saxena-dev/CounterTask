package routes

import (
	"goPractice/controllers"

	"github.com/gin-gonic/gin"
)

type TaskRoutes interface {
	RegisterHandlers()
}

type taskRoutes struct {
	engine     *gin.Engine
	controller controllers.TaskController
}

func NewTaskRoutes(engine *gin.Engine, controller controllers.TaskController) TaskRoutes {

	return &taskRoutes{
		engine:     engine,
		controller: controller,
	}
}

func (taskRoutes *taskRoutes) RegisterHandlers() {
	taskRoutes.engine.POST("/createTaskCounter", taskRoutes.controller.CreateTaskCounter)
	taskRoutes.engine.GET("/getCounterProgress/:taskId", taskRoutes.controller.GetTaskProgress)
	taskRoutes.engine.DELETE("/removeCounterTask/:taskId", taskRoutes.controller.RemoveCounterTask)

}
