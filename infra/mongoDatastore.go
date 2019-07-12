package datastore

import (
	"context"
	"sync"
	"time"

	"bitbucket.org/libertywireless/wonderwall-auth/wlog"

	"bitbucket.org/libertywireless/wonderwall-auth/config"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
	"github.com/sirupsen/logrus"
)

const CONNECTED = "Successfully connected to database: %v"

type MongoDatastore struct {
	DB      *mongo.Database
	Session *mongo.Client
	logger  *logrus.Logger
}

func NewDatastore(config config.GeneralConfig, logger *logrus.Logger) *MongoDatastore {

	var mongoDataStore *MongoDatastore
	db, session := connect(config, logger)
	if db != nil && session != nil {

		// log statements here as well

		mongoDataStore = new(MongoDatastore)
		mongoDataStore.DB = db
		mongoDataStore.logger = logger
		mongoDataStore.Session = session
		return mongoDataStore
	}

	logger.Fatalf("Failed to connect to database: %v", config.DatabaseName)

	return nil
}

func connect(generalConfig config.GeneralConfig, logger *logrus.Logger) (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		db, session = connectToMongo(generalConfig, logger)
	})

	return db, session
}

func connectToMongo(generalConfig config.GeneralConfig, logger *logrus.Logger) (a *mongo.Database, b *mongo.Client) {

	var err error
	session, err := mongo.NewClient(generalConfig.DatabaseHost)
	if err != nil {
		logger.Fatal(err)
	}
	session.Connect(context.TODO())
	if err != nil {
		logger.Fatal(err)
	}

	var DB = session.Database(generalConfig.DatabaseName)
	logger.Info(CONNECTED, generalConfig.DatabaseName)

	return DB, session
}

func PopulateIndex(store *MongoDatastore, collection string) {
	c := store.DB.Collection(collection)
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	index := yieldIndexModel()
	_, err := c.Indexes().CreateOne(context.Background(), index, opts)
	if err == nil {
		wlog.Logger.Println("Successfully created the index")
	} else {
		wlog.Logger.Errorln("Could not create index")
	}
}

func yieldIndexModel() mongo.IndexModel {
	keys := bsonx.Doc{{Key: "email", Value: bsonx.Int32(int32(1))}}
	index := mongo.IndexModel{}
	index.Keys = keys
	t := true
	index.Options = &options.IndexOptions{
		Unique: &t}
	return index
}
