package mysqldb

import (
	"database/sql"
	"databaseConnector/dto"
	"databaseConnector/models"
	"fmt"
	"log"
)

// CreateCustomer -> create a new customer by received auth dto
func (this MySqlDb) CreateCustomer(dto *dto.SignInDto) error {

	// use IsEntityExists before saving
	sql := "INSERT INTO Customer (name, email, password) VALUES (?,?,?)"
	defer this.Disconnect()

	result := this.db.QueryRow(sql, &dto.Name, &dto.Email, &dto.Password)
	return result.Err()
}

// IsCustomerExists -> check is db has a customer with received username
func (this MySqlDb) IsEntityExists(email *string) (bool, error) {

	var id uint64
	sql := "SELECT id FROM Customer WHERE email=?"
	defer this.Disconnect()

	rows, err := this.db.Query(sql, email)
	if err != nil {
		log.Fatal("selection error", err)
	}
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Println("scan error -> ", err)
			return false, err
		}
	}

	return true, nil
}

// FindById -> find cuxtomer by received id
func (this MySqlDb) FindById(id *uint32) (*models.Customer, error) {

	var c models.Customer
	defer this.Disconnect()

	rows, err := this.db.Query("SELECT * FROM Customer WHERE id=?", &id)
	if err != nil {
		log.Println("selection error", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&c.Id, &c.Name, &c.Email, &c.Password, &c.CreatedAt)
		if err != nil {
			log.Println("scan error -> ", err)
			return nil, err
		}
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(err.Error())
			// return c, fmt.Errorf("customer %d: no such customer", id)
		}
		return nil, err
	}

	return &c, err
}
