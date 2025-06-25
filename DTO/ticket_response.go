package DTO

import "github.com/google/uuid"

type Response struct {
	TicketID uuid.UUID
	UserID   string
	Message  string
}
