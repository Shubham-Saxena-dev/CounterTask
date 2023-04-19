package main

import (
	"database/sql"
	"fmt"
	"goPractice/controller"
	"goPractice/database"
	errorhandler "goPractice/errorHandler"
	"goPractice/repository"
	"goPractice/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router       *gin.Engine
	dbInstance   *sql.DB
	errorHandler errorhandler.ErrorHandler
)

func main() {

	fmt.Println("Create task project")
	initDatabase()
	createServer()

}

func createServer() {
	router = gin.Default()
	intializeLayers()
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
	repository = repository.NewDataBase(*dbInstance, errorHandler)
	service = service.NewTaskService(repository, errorHandler)
	controller = controller.NewController(service, errorHandler)
}
