package database

import (
	"fmt"
	"log"
	"ticketing_server/entity"

	"gorm.io/driver/postgres" // GORM PostgreSQL driver
	"gorm.io/gorm"
)

var GormDB *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "mockticket"
)

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Kolkata",
		host, user, password, dbname, port)

	var err error
	GormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	sqlDB, err := GormDB.DB()
	if err != nil {
		log.Printf("Failed to get underlying sql.DB: %v", err)
		return err
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Printf("Failed to ping database: %v", err)
		return err
	}

	fmt.Println("Successfully connected to database with GORM!")

	err = GormDB.AutoMigrate(&entity.User{}, &entity.Ticket{})
	if err != nil {
		log.Printf("Failed to auto migrate database schema: %v", err)
		return err
	}
	fmt.Println("Database tables auto-migrated successfully.")
	return nil
}

func GetGormDB() *gorm.DB {
	if GormDB == nil {
		log.Fatal("Database not initialized. Call InitDB() first!")
	}
	return GormDB
}
