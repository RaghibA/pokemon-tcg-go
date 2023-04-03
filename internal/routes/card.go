package routes

import (
	"github.com/RaghibA/pokemon-tcg-go/internal/handlers"
	"github.com/RaghibA/pokemon-tcg-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Card(r *gin.Engine) {
	cardRouter := r.Group("/card")
	{
		cardRouter.GET("/query", handlers.QueryCardHandler)
		cardRouter.POST("/add", middleware.Authenticate, handlers.AddCardHandler)
		cardRouter.GET("/getAll", middleware.Authenticate, handlers.GetCardsHandler)
		cardRouter.DELETE("/remove", middleware.Authenticate, handlers.DeleteCardHandler)
	}
}
