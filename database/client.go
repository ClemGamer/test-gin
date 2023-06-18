package database

import (
	"log"

	"github.com/ClemGamer/test-gin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

// connect to sql server
func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Can not connect to DB")
	}
	log.Println("Connected to Database")
}

// migrate tables
func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed")
}

// close sql connection
func Close() {
	sqlDB, err := Instance.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
	log.Println("sql closed")
}
