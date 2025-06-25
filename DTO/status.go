package DTO

import (
	"database/sql/driver"
	"fmt"
)

type TicketStatus int

const (
	RAISED         TicketStatus = iota // 0
	IN_PROGRESS                        // 1
	COMPLETED                          // 2
	UNKNOWN_STATUS                     // 3
)

func (ts TicketStatus) String() string {
	switch ts {
	case RAISED:
		return "Raised"
	case IN_PROGRESS:
		return "In Progress"
	case COMPLETED:
		return "Completed"
	default:
		return fmt.Sprintf("Unknown Status (%d)", ts)
	}
}

func (ts TicketStatus) Value() (driver.Value, error) {
	return ts.String(), nil
}

// Scan implements the sql.Scanner interface.
// This method is called by GORM when loading a TicketStatus from the database.
func (ts *TicketStatus) Scan(value interface{}) error {
	if value == nil {
		*ts = UNKNOWN_STATUS // Set to UNKNOWN_STATUS if the DB value is NULL
		return nil
	}

	var s string
	switch v := value.(type) {
	case []byte:
		s = string(v)
	case string:
		s = v
	default:
		return fmt.Errorf("unsupported type for TicketStatus scan: %T", value)
	}

	switch s {
	case "Raised":
		*ts = RAISED
	case "In Progress":
		*ts = IN_PROGRESS
	case "Completed":
		*ts = COMPLETED
	default:
		*ts = UNKNOWN_STATUS
		return fmt.Errorf("unknown TicketStatus string from DB: %s", s)
	}
	return nil
}
