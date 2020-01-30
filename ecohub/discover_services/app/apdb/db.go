package apdb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"github.com/discover_services/config"
	"github.com/discover_services/logger"
	_"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	_"github.com/discover_services/app/model"
	_"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"fmt"
	_"reflect"
	_"errors"

)


const DbName = "runner"
var logs *logger.Logger
//var conf config.Config
var MongoClient *mongo.Client

func Init() {
	logs = logger.Logs.GetLogger("apdb")
	if config.CONF.Database.Dialect == "mongodb" {
		MongodbInitialize()
	}
	
}

func MongodbInitialize() {	
	dbURI := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
	config.CONF.Database.Dialect,
	config.CONF.Database.Username,
	config.CONF.Database.Password,
	config.CONF.Database.Host,
	config.CONF.Database.Port,
	config.CONF.Database.Name)
	clientOptions := options.Client().ApplyURI(dbURI)

	//Create connections 
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logs.Fatalf("Could not connect mongodb, Error:%s", err)
	}

	//Check connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logs.Fatalf("Could not connect to mongodb, Error:%s", err)
	}
	logs.Infof("Sucessfully Connected to MongoDB!")
	MongoClient = client
}


type Database interface {	
	

}

type Mongodb struct {
    Client *mongo.Client
}

var mongodb Mongodb

func GetDatabase() Database {
	if config.CONF.Database.Dialect == "mongodb" {
		mongodb.Client = MongoClient
		return &mongodb
	}
	return &mongodb
	
}




func PopulateIndex(c *mongo.Collection) {
	//c := obj.DbClient.Database(obj.DbName).Collection(obj.CollectionName)
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	index := yieldIndexModel()
	c.Indexes().CreateOne(context.Background(), index, opts)
	//logs.Infof("Successfully create the index")
}

func yieldIndexModel() mongo.IndexModel {
	keys := bsonx.Doc{{Key: "name", Value: bsonx.String("text")}}
	index := mongo.IndexModel{}
	index.Keys = keys
	index.Options = options.Index().SetUnique(true) //bsonx.Doc{{Key: "unique", Value: bsonx.Boolean(true)}}
	return index
}


