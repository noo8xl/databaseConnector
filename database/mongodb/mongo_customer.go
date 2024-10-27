package mongodb

import (
	"context"
	"databaseConnector/dto"
	"databaseConnector/models"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateCustomer -> create a new customer by received auth dto
func (this MongoDb) CreateCustomer(dto *dto.SignInDto) error {

	defer this.Disconnect()

	db := this.db.Database("test_db")
	ctx := context.TODO()
	collection := db.Collection("Customer")

	nextId, err := getNextID(ctx, collection)
	if err != nil {
		log.Println("getNextID func returns an error: ", err)
		return err
	}

	c := models.Customer{
		Id:        nextId,
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  dto.Password,
		CreatedAt: time.Now(),
	}

	_, err = collection.InsertOne(ctx, &c)
	if err != nil {
		fmt.Println("insertion err: ", err)
		return err
	}

	return err
}

// IsCustomerExists -> check is db has a customer with received username
func (this MongoDb) IsEntityExists(email *string) (bool, error) {

	var customer models.Customer
	filter := bson.D{{Key: "email", Value: email}}
	defer this.Disconnect()

	db := this.db.Database("test_db")
	ctx := context.TODO()
	collection := db.Collection("Customer")
	opts := options.FindOneOptions{}

	result := collection.FindOne(ctx, filter, &opts)
	if result.Err() != nil {
		fmt.Println("IsCustomerExists got an err: ", result.Err())
		return false, result.Err()
	}

	err := result.Decode(&customer)
	if err != nil {
		fmt.Println("IsCustomerExists got an err at decoding: ", result.Err())
		return false, err
	}

	return true, nil
}

// FindById -> find cuxtomer by received id
func (this MongoDb) FindById(id *uint32) (*models.Customer, error) {

	var err error
	var customer models.Customer
	filter := bson.D{{Key: "_id", Value: id}}
	defer this.Disconnect()

	db := this.db.Database("test_db")
	ctx := context.TODO()
	collection := db.Collection("Customer")
	opts := options.FindOneOptions{}

	result := collection.FindOne(ctx, filter, &opts)
	if result != nil {
		result.Decode(&customer)
	}

	return &customer, err
}
