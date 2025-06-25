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
	dbname   = "mockTicket"
)

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Kolkata",
		host, user, password, dbname, port)

	var err error
	GormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := GormDB.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying sql.DB: %v", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to database with GORM!")

	err = GormDB.AutoMigrate(&entity.User{}, &entity.Ticket{}) // <--- MODIFIED HERE
	if err != nil {
		log.Fatalf("Failed to auto migrate database schema: %v", err)
	}
	fmt.Println("Database tables auto-migrated successfully.")
}

// GetGormDB returns the GORM database instance.
func GetGormDB() *gorm.DB {
	return GormDB
}
