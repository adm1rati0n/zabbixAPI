package dbConnection

import (
	"configHandler"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"models"
)

var DB *sql.DB
var Config models.ConfigModel

func DBconnect() {
	connectionString := configHandler.GetConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		connectionString.DbHost, connectionString.Port, connectionString.User, connectionString.Password, connectionString.DbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")
	Config = connectionString
	DB = db
}
