package routes

import (
	"github.com/RaghibA/pokemon-tcg-go/internal/handlers"
	"github.com/RaghibA/pokemon-tcg-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		// Create Account
		userRouter.POST("/register", handlers.CreateUserHandler)
		// Login
		userRouter.POST("/login", handlers.LoginUserHandler)
		// Delete
		userRouter.DELETE("/deactivate", middleware.Authenticate, handlers.DeleteUserHandler)
		// Authenticate
		userRouter.POST("/authenticate", middleware.Authenticate, handlers.AuthenticateUserHandler)
	}
}
