package entity

import (
	"github.com/google/uuid"
)

// User represents the 'users' table
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	TicketIDs []int     `gorm:"type:integer[]"`
	Tickets   []Ticket  `gorm:"foreignKey:UserID"`
}
