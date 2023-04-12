package routes

import (
	"github.com/RaghibA/pokemon-tcg-go/internal/handlers"
	"github.com/RaghibA/pokemon-tcg-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", handlers.CreateUserHandler)
		userRouter.POST("/login", handlers.LoginUserHandler)
		userRouter.DELETE("/deactivate", middleware.Authenticate, handlers.DeleteUserHandler)
		userRouter.POST("/authenticate", middleware.Authenticate, handlers.AuthenticateUserHandler)
		userRouter.POST("/logout", middleware.Authenticate, handlers.LogoutUserHandler)
	}
}
