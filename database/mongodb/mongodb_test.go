package mongodb

import (
	"databaseConnector/dto"
	"testing"
)

func TestInitDbConnection(t *testing.T) {
	db, err := InitDbConnection()
	if err != nil {
		t.Errorf("got an error at InitDbConnection: %q", err)
	}
	defer db.Disconnect()
}

func TestCreateCustomer(t *testing.T) {
	dto := dto.SignInDto{
		Name:     "test1",
		Email:    "test@test.com",
		Password: "testPwd123",
	}

	db, err := InitDbConnection()
	if err != nil {
		t.Errorf("got an error at InitDbConnection: %q", err)
	}

	err = db.CreateCustomer(&dto)
	if err != nil {
		t.Errorf("got an error at CreateCustomer: %q", err)
	}
}

func TestFindById(t *testing.T) {
	var id uint32 = 1

	db, err := InitDbConnection()
	if err != nil {
		t.Errorf("got an error at InitDbConnection: %q", err)
	}

	_, err = db.FindById(&id)
	if err != nil {
		t.Errorf("got an error at FindById: %q", err)
	}
}

func TestIsEntityExists(t *testing.T) {
	db, err := InitDbConnection()
	if err != nil {
		t.Errorf("got an error at InitDbConnection: %q", err)
	}

	var email string = "test@test.com"
	_, err = db.IsEntityExists(&email)
	if err != nil {
		t.Errorf("got an error at IsEntityExists: %q", err)
	}
}
