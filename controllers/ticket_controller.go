package controllers

import (
	"log"
	"net/http"
	"ticketing_server/DTO"
	"ticketing_server/database"
	"ticketing_server/entity"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	validate *validator.Validate = validator.New()
)

func RaiseTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		var db *gorm.DB = database.GetGormDB()
		var (
			err     error
			request DTO.RaiseRequest
		)

		if err = c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Cannot Bind Request to internal request struct",
				"error":   err.Error(),
			})
			return
		}

		if err := validate.Struct(request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Validator Error",
				"error":   err.Error(),
			})
			return
		}

		parsedUserID, err := uuid.Parse(request.UserID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID format. Must be a valid UUID."})
			return
		}

		var user entity.User
		user.ID = parsedUserID
		result := db.FirstOrCreate(&user, entity.User{ID: parsedUserID})
		if result.Error != nil {
			log.Printf("Error finding or creating user %s: %v", parsedUserID, result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find or create user."})
			return
		}

		newTicket := entity.Ticket{
			Description: request.Issue,
			Status:      DTO.RAISED,
			Title:       request.Title,
			CreatedAt:   time.Now(),
			UserID:      user.ID, // Link to the (found or created) user's ID
		}

		if result := db.Create(&newTicket); result.Error != nil {
			log.Printf("Failed to create ticket for user %s: %v", user.ID, result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ticket."})
			return
		}

		res := DTO.RaiseResponse{
			TicketID:  newTicket.ID,
			UserID:    newTicket.UserID,
			Message:   "Ticket raised successfully.",
			CreatedAt: newTicket.CreatedAt,
			Status:    DTO.RAISED.String(),
		}

		c.JSON(http.StatusCreated, res)
	}
}

func GetTicket() gin.HandlerFunc {
	return func(c *gin.Context) {
		var db *gorm.DB = database.GetGormDB()
		var (
			err     error
			request DTO.GetRequest
		)

		if err = c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Cannot Bind Request to internal request struct",
				"error":   err.Error(),
			})
			return
		}

		parsedUserID, err := uuid.Parse(request.UserID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID format. Must be a valid UUID."})
			return
		}
		var user entity.User
		userResult := db.First(&user, "id = ?", parsedUserID)

		if userResult.Error != nil {
			if userResult.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusOK, gin.H{
					"message": "You have no ongoing tickets (User has not created any tickets).",
					"tickets": []DTO.TicketSummary{},
				})
				return
			}
			log.Printf("Error checking user existence for %s: %v", parsedUserID, userResult.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error checking user existence."})
			return
		}

		var tickets []entity.Ticket
		ticketsResult := db.Where("user_id = ?", parsedUserID).Find(&tickets)
		if ticketsResult.Error != nil {
			log.Printf("Error fetching tickets for user %s: %v", parsedUserID, ticketsResult.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tickets for the user."})
			return
		}

		var ticketSummaries []DTO.TicketSummary
		if len(tickets) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "User exists but has no ongoing tickets.",
				"tickets": []DTO.TicketSummary{},
			})
			return
		}

		for _, ticket := range tickets {
			ticketSummaries = append(ticketSummaries, DTO.TicketSummary{
				TicketID: ticket.ID,
				Title:    ticket.Title,
				Status:   ticket.Status.String(),
			})
		}

		res := DTO.GetTicketsResponse{
			Tickets: ticketSummaries,
			Message: "Tickets retrieved successfully.",
		}
		c.JSON(http.StatusOK, res)
	}
}
