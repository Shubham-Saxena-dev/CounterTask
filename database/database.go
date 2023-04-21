package database

import (
	"database/sql"
	"goPractice/errorhandler"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

const (
	ConnectionString = "root:admin@/myDb"
	Type             = "mysql"
)

var (
	dbInstance *sql.DB
	once       sync.Once
)

func GetDatabaseConnection(errorHandler errorhandler.ErrorHandler) *sql.DB {

	if dbInstance == nil {
		once.Do(
			func() {
				dbInstance = newDatabaseConnection(errorHandler)
			})
	}
	return dbInstance
}

func newDatabaseConnection(errorHandler errorhandler.ErrorHandler) *sql.DB {

	db, err := sql.Open(Type, ConnectionString)

	if err != nil {
		log.Error(err)
	}

	err = db.Ping()
	if err != nil {
		log.Error(err)
	}

	db.Query("DROP table Tasks")
	_, err = db.Query("CREATE TABLE Tasks(taskId int, start int, end int);")

	if err != nil {
		log.Error(err)

	}

	log.Info("Connected to database")

	return db
}
