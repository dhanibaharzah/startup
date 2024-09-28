package router

import (
	"startup/internal/auth"
	"startup/internal/user"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Authentication routes
	r.POST("/register", user.Register)
	r.POST("/login", user.Login)
	r.POST("/logout", user.Logout)

	// Protected routes
	authorized := r.Group("/users")
	authorized.Use(auth.Middleware())
	{
		authorized.GET("/", user.ListUsers)
		authorized.GET("/:id", user.GetUser)
		authorized.POST("/", auth.AdminMiddleware(), user.HandleCreateUser)
		authorized.PUT("/:id", auth.AdminMiddleware(), user.HandleUpdateUser)
		authorized.DELETE("/:id", auth.SuperAdminMiddleware(), user.HandleDeleteUser)
	}

	return r
}
