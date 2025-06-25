package DTO

import (
	"time"

	"github.com/google/uuid"
)

type RaiseResponse struct {
	TicketID  uint      `json:"ticket_id"`
	UserID    uuid.UUID `json:"user_id"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type TicketSummary struct {
	TicketID uint   `json:"ticket_id"`
	Title    string `json:"title"`
	Status   string `json:"status"`
}

type GetTicketsResponse struct {
	Tickets []TicketSummary `json:"tickets"`
	Message string          `json:"message"`
}
