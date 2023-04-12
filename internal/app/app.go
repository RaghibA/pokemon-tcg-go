package app

import (
	"log"
	"os"

	"github.com/RaghibA/pokemon-tcg-go/internal/database"
	"github.com/RaghibA/pokemon-tcg-go/internal/routes"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Run() {
	log.Println("Starting server...")

	// Connect to DB
	database.Connect()

	if os.Getenv("DEPLOY_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./public/assets", true)))

	// Register routers
	routes.User(r)
	routes.Card(r)

	// r.Static("/", "../../client/build")
	// url := fmt.Sprintf("%s:%s", os.Getenv("API_URL"), os.Getenv("API_PORT"))

	r.Run(":8080")
}
