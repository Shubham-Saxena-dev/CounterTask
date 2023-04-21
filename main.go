package main

import (
	"database/sql"
	"fmt"
	"goPractice/controllers"
	"goPractice/database"
	"goPractice/errorhandler"
	"goPractice/repository"
	"goPractice/routes"
	"goPractice/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	dbInstance   *sql.DB
	errorHandler errorhandler.ErrorHandler
	repo         repository.TaskRepository
	serv         service.TaskService
	controller   controllers.TaskController
)

func main() {

	fmt.Println("Create task project")
	initDatabase()
	createServer()

}

func createServer() {
	router := gin.Default()
	intializeLayers()
	routes.NewTaskRoutes(router, controller)
	err := router.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func initDatabase() {
	log.Info("Connecting to MySQL")
	createErrorHandler()
	dbInstance = database.GetDatabaseConnection(errorHandler)
}

func createErrorHandler() {
	errorHandler = errorhandler.NewErrorHandler()
}

func intializeLayers() {
	repo = repository.NewDataBase(*dbInstance, errorHandler)
	serv = service.NewTaskService(repo, errorHandler)
	controller = controllers.NewTaskController(serv, errorHandler)
}
