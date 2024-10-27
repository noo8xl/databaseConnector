package database

import (
	"databaseConnector/database/mysqldb"
	"databaseConnector/dto"
	models "databaseConnector/models"
	"log"
)

type dbcontract interface {
	Disconnect()
	IsEntityExists(key *string) (bool, error)

	CreateCustomer(dto *dto.SignInDto) error
	FindById(id *uint32) (*models.Customer, error)
}

type Database struct {
	db dbcontract
}

func NewDatabase(db dbcontract) *Database {
	return &Database{db: db}
}

func init() {
	_, err := mysqldb.InitDbConnection()
	if err != nil {
		log.Fatal("mysql init err", err)
	}

	// var id uint32 = 1
	// database := NewDatabase(db)
	// database.db.FindById(&id)
}
