package repository

import (
	"context"
	"goPractice/errorhandler"
	"goPractice/models"
	"goPractice/utils"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	taskMap = make(map[string]utils.Progress)
)

type TaskRepository interface {
	SaveTaskCounters(request models.TaskCounterRequest) string
	RemoveCounterTask(taskId string) string
	GetTaskProgress(taskId string) int
}

type taskRepository struct {
	mongo        *mongo.Collection
	context      context.Context
	errorHandler errorhandler.ErrorHandler
}

func NewDataBase(db mongo.Collection, ctx context.Context, errorHandler errorhandler.ErrorHandler) TaskRepository {

	return &taskRepository{
		mongo:        &db,
		context:      ctx,
		errorHandler: errorHandler,
	}
}

func (db *taskRepository) SaveTaskCounters(request models.TaskCounterRequest) string {
	result, err := db.mongo.InsertOne(db.context, request)
	if err != nil {
		db.errorHandler.HandleError(err)
	}
	log.Info("Created in db with id: %s", result.InsertedID)

	taskUtil := utils.NewProgressTask(request)

	taskMap[request.TaskId] = taskUtil
	go func() {
		taskUtil.Calculate()
	}()
	return request.TaskId
}

func (db *taskRepository) RemoveCounterTask(taskId string) string {
	value, ok := taskMap[taskId]
	if !ok {
		return "NOT_FOUND"
	}
	value.CancelTask()

	if value.GetStatus() == "Cancelled" {
		delete(taskMap, taskId)
		db.mongo.DeleteOne(context.TODO(), bson.M{"TaskId": taskId})
	}

	return "Task Cancelled"
}

func (db *taskRepository) GetTaskProgress(taskId string) int {
	value, ok := taskMap[taskId]
	if !ok {
		return -1
	}
	progress := value.GetProgress()
	if progress > 0 {
		return progress
	}
	return 0
}
