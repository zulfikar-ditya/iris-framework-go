package config_database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDatabaseCredentials() (string, string, string, string, string) {
	DATABASE_HOST := os.Getenv("DATABASE_HOST")
	DATABASE_PORT := os.Getenv("DATABASE_PORT")
	DATABASE_NAME := os.Getenv("DATABASE_NAME")
	DATABASE_USER := os.Getenv("DATABASE_USER")
	DATABASE_PASSWORD := os.Getenv("DATABASE_PASSWORD")

	return DATABASE_HOST, DATABASE_PORT, DATABASE_NAME, DATABASE_USER, DATABASE_PASSWORD
}

func Connect() (*gorm.DB, error) {
	DATABASE_HOST, DATABASE_PORT, DATABASE_NAME, DATABASE_USER, DATABASE_PASSWORD := getDatabaseCredentials()

	// "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=" + DATABASE_HOST + " user=" + DATABASE_USER + " password=" + DATABASE_PASSWORD + " dbname="+ DATABASE_NAME +" port=" + DATABASE_PORT + " sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	return db, nil
}