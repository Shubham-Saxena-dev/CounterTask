package database

import (
	"context"
	"goPractice/errorhandler"
	"sync"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ConnectionString = "mongodb://localhost:27017"
	Type             = "mysql"
)

var (
	dbInstance *mongo.Collection
	once       sync.Once
)

func GetDatabaseConnection(errorHandler errorhandler.ErrorHandler) *mongo.Collection {

	if dbInstance == nil {
		once.Do(
			func() {
				dbInstance = newDatabaseConnection(errorHandler)
			})
	}
	return dbInstance
}

func newDatabaseConnection(errorHandler errorhandler.ErrorHandler) *mongo.Collection {

	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(ConnectionString))

	if err != nil {
		log.Error(err)
		panic(err)
	}

	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	usersCollection := db.Database("tasks").Collection("counter")

	log.Info("Connected to database")

	return usersCollection
}
