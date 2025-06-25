package routes

import (
	"ticketing_server/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incoming_routes *gin.Engine) {
	incoming_routes.POST("/raiseticket", controllers.GetTicket())
	incoming_routes.GET("/getticket", controllers.GetTicket())
}
