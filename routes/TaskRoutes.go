package routes

import (
	"goPractice/controller"

	"github.com/gin-gonic/gin"
)

type TaskRoutes interface {
	RegisterHandlers(server *gin.Engine, controller controller.Controller)
}

type taskRoutes struct {
	server     *gin.Engine
	controller controller.Controller
}

func NewTaskRoutes(server *gin.Engine, controller controller.Controller) TaskRoutes {

	return &taskRoutes{
		server:     server,
		controller: controller,
	}
}

func (taskRoutes *taskRoutes) RegisterHandlers(server *gin.Engine, controller controller.Controller) {

}
