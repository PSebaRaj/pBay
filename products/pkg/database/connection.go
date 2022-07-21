package database

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // global connection to DB, only use within DB package
var err error

// connect to DB using GORM
func Connect(dialect, host, dbport, user, dbname, dbpassword string) *sql.DB {

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, dbpassword, host, dbport, dbname)

	// Opening connection to database
	DB, err = gorm.Open(mysql.Open(mysqlURI), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the database")
		panic(err)
	}

	DB_Interface, err := DB.DB()
	if err != nil {
		fmt.Println("Could not convert the DB into a generic SQL interface")
		panic(err)
	}

	fmt.Printf("Connected to %s DB @ %s:%s\n", dialect, host, dbport)

	return DB_Interface
}

