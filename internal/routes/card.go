package routes

import "github.com/gin-gonic/gin"

func Card(r *gin.Engine) {
	cardRouter := r.Group("/card")
	{
		cardRouter.POST("/query")
	}
}
