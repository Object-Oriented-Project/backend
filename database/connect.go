package database

import (
	"fmt"
	"log"

	"github.com/Object-Oriented-Project/backend/config"
	"github.com/Object-Oriented-Project/backend/internals/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	host := config.Config("DB_HOST")
	port := config.Config("DB_PORT")
	user := config.Config("DB_USER")
	password := config.Config("DB_PASSWORD")
	dbname := config.Config("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Database connected")

	// Migrate the schema
	DB.AutoMigrate(&model.User{}, &model.Menu{})
	fmt.Println("Database migration completed!")
}