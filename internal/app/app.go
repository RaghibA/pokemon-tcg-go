package app

import (
	"fmt"
	"log"
	"os"

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
	routes.Card(r)

	fmt.Sprintf("%s:%s", os.Getenv("API_URL"), os.Getenv("API_PORT"))
	r.Run(":8080")
}
