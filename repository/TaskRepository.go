package repository

import (
	"database/sql"
	"goPractice/errorhandler"
)

type TaskRepository interface {
	SaveTaskCounters(x, y, taskId int) (bool, error)
	RemoveCounterTask(taskId int) (bool, error)
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

func (db *taskRepository) SaveTaskCounters(x, y, taskId int) (bool, error) {
	return true, nil
}

func (db *taskRepository) RemoveCounterTask(taskId int) (bool, error) {
	return true, nil
}
