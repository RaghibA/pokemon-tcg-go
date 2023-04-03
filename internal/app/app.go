package app

import (
	"log"

	"github.com/RaghibA/pokemon-tcg-go/internal/database"
	"github.com/RaghibA/pokemon-tcg-go/internal/routes"
	"github.com/gin-gonic/gin"
)

func Run() {
	log.Println("Starting server...")

	// Connect to DB
	database.Connect()

	r := gin.Default()

	// Register routers
	routes.User(r)

	r.Run(":8080")
}
