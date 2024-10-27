package mysqldb

import (
	"database/sql"
	"databaseConnector/config"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlDb struct {
	db *sql.DB
}

// InitDbConnection -> initialize a database connection with recieved options
func InitDbConnection() (*MySqlDb, error) {

	dbDriver := "mysql"
	opts := config.MySqlOptions()
	str := strings.Join([]string{opts.User, ":", opts.Pwd, "@tcp(", opts.Host, ":", opts.Port, ")/", opts.Name, "?parseTime=true"}, "")
	db, err := sql.Open(dbDriver, str)
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(40)
	return &MySqlDb{db: db}, err
}

// Disconnect -> disconnect from db after a func will be done
func (this MySqlDb) Disconnect() {
	err := this.db.Close()
	if err != nil {
		log.Fatalf("Mysql close falure: %v", err)
	}
}
