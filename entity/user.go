package entity

import (
	"github.com/google/uuid"
)

// User represents the 'users' table
type User struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Tickets []Ticket  `gorm:"foreignKey:UserID"`
}
