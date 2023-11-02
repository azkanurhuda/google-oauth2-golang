package initializers

import (
	"fmt"
	"github.com/azkanurhuda/google-oauth2-golang/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("golang.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to Database")
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed migration")
	}
	fmt.Println("Connected Successfully to the Database")
}
