package config

import (
	"fmt"
	"log"
	"os"

	"github.com/itsharshitk/1_ToDoCRUD/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB() error {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error in Loading .env File")
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		return err
	}

	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatal("Failed to Get SQL DB: ", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Database Ping Issue: ", err)
	}

	fmt.Println("Database Connected Successfully")

	// if err := Db.AutoMigrate(&model.User{}); err != nil {
	// 	log.Fatalf("Failed to Automigrate User: %v", err)
	// 	return err
	// }

	allModels := []any{
		&model.User{},
		&model.Todo{},
	}

	if err := Db.AutoMigrate(allModels...); err != nil {
		log.Fatalf("Failed to Automigrate: %v", err)
		return err
	}

	return nil
}
