package main

import (
	"task_management_system/internal/db"
	"task_management_system/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	r := gin.Default()

	// Public routes
	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	// Protected routes
	protected := r.Group("/")
	protected.Use(handlers.AuthMiddleware())
	{
		protected.GET("/tasks", handlers.GetTasks)
	}

	r.Run(":8080")
}
