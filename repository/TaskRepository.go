package repository

import (
	"context"
	"goPractice/errorhandler"
	"goPractice/models"

	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	SaveTaskCounters(request models.TaskCounterRequest) uuid.UUID
	RemoveCounterTask(taskId int) (bool, error)
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

func (db *taskRepository) SaveTaskCounters(request models.TaskCounterRequest) uuid.UUID {
	result, err := db.mongo.InsertOne(db.context, request)
	if err != nil {
		db.errorHandler.HandleError(err)
	}
	log.Info("Created in db with id: %s", result.InsertedID)

	return request.TaskId
}

func (db *taskRepository) RemoveCounterTask(taskId int) (bool, error) {


	return true, nil
}
