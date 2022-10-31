package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	err := godotenv.Load("./.env")
	if err != nil {
		println("Gagal terkoneksi ke database, ENV Error")
		return nil
	}
	dsn := os.Getenv("DSN")

	fmt.Println("Connecting to database...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		println("Gagal terkoneksi ke database")
		return nil
	}
	fmt.Println("Connected")
	return db
}
