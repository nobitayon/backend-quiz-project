package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GormDB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("yon:a@/quizproject?parseTime=true"), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		panic("could not connect to the database")
	}

	GormDB = connection
}
