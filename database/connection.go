package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GormDB *gorm.DB

func Connect() {

	username := os.Getenv("username")
	password := os.Getenv("password")
	database := os.Getenv("database")

	connection, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@/%s?parseTime=true", username, password, database)), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		panic("could not connect to the database")
	}

	GormDB = connection
}
