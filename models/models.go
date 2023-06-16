package models

import (
	"context"
	"hito/configs"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DbName, mongoUrl string
var MongoClient *mongo.Client

func init() {
	mongoUrl = configs.GeneralConf.GetString("mongo.url")
	DbName = configs.GeneralConf.GetString("mongo.dbName")
}

func Init() {
	// Create a new client and connect to the server
	log.Infof("mongo url is %s", mongoUrl)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUrl))
	if err != nil {
		panic(err)
	}
	MongoClient = client

	// Ping the primary
	if err := MongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	log.Info("Successfully connected mongodb and pinged.")
}
