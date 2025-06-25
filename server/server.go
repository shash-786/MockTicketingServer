package server

import (
	"log"
	"os"
	"ticketing_server/database"
	"ticketing_server/routes"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := database.InitDB()
	if err != nil {
		log.Fatalf("database init error")
	}

	router := gin.Default()
	routes.UserRoutes(router)
	log.Fatal(router.Run(":" + port))
}
