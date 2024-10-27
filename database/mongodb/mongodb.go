package mongodb

import (
	"context"
	"databaseConnector/config"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDb struct {
	db *mongo.Client
}

// InitDbConnection -> initialize a database connection with recieved options
func InitDbConnection() (*MongoDb, error) {

	ctx := context.Background()
	opts := config.MongoDbOptions()
	path := fmt.Sprintf("mongodb+srv://%s:%s@cluster001.sipjs.mongodb.net/?retryWrites=true&w=majority", opts.User, opts.Pwd)
	clientOptions := options.Client().ApplyURI(path)

	db, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoDb{db: db}, err
}

// Disconnect -> disconnect from db after a func will be done
func (this MongoDb) Disconnect() {
	err := this.db.Disconnect(context.Background())
	if err != nil {
		log.Fatalf("MongoDb close falure: %v", err)
	}
}

// getNextID -> this func using for the auto incrementing int id value
func getNextID(ctx context.Context, collection *mongo.Collection) (uint32, error) {

	var counter struct {
		Id uint32 `bson:"_id"`
	}

	filter := bson.D{}
	opts := *options.FindOne().SetSort(bson.D{{Key: "_id", Value: -1}})

	err := collection.FindOne(ctx, filter, &opts).Decode(&counter)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			log.Println("got a not found exception -> ", err)
			return 1, nil
		}
		log.Println("got an error -> ", err)
		return 0, err
	}

	return counter.Id + 1, nil
}
