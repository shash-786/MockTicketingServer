package DTO

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	TicketID  uuid.UUID
	UserID    string
	Message   string
	CreatedAt time.Time
}
