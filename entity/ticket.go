package entity

import (
	"ticketing_server/DTO"
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID          uint             `gorm:"primaryKey;autoIncrement"`
	Description string           `gorm:"type:text"`
	Status      DTO.TicketStatus `gorm:"type:text"`
	Title       string           `gorm:"type:text"`
	CreatedAt   time.Time
	UserID      uuid.UUID `gorm:"type:uuid"`
}
