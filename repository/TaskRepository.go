package repository

import (
	"database/sql"
	errorhandler "goPractice/errorHandler"
)

type TaskRepository interface {
	saveTaskCounters(x, y, taskId int) (bool, error)
	removeCounterTask(taskId int) (bool, error)
}

type taskRepository struct {
	db           *sql.DB
	errorHandler errorhandler.ErrorHandler
}

func NewDataBase(db sql.DB, errorHandler errorhandler.ErrorHandler) TaskRepository {

	return &taskRepository{
		db:           &db,
		errorHandler: errorHandler,
	}
}

func (db *taskRepository) saveTaskCounters(x, y, taskId int) (bool, error) {
	return true, nil
}

func (db *taskRepository) removeCounterTask(taskId int) (bool, error) {
	return true, nil
}
