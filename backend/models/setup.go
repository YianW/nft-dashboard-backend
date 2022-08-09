package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB() {
	_ = godotenv.Load(".env")

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Hong_Kong", DbHost, DbUser, DbPassword, DbName, DbPort)

	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database ")
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to database ")
	}
	err = Db.AutoMigrate(&User{})
	err = Db.AutoMigrate(&Transaction{})
	err = Db.AutoMigrate(&NFT{})
	err = Db.AutoMigrate(&ManualTransfer{})
	err = Db.AutoMigrate(&NFTMintHistory{})
	err = Db.AutoMigrate(&Contract{})
	err = Db.AutoMigrate(&Collection{})
	err = Db.AutoMigrate(&APIkey{})
	err = Db.AutoMigrate(&MetaInfoTemplate{})
	// TODO : AutoMigrate error handling
}
