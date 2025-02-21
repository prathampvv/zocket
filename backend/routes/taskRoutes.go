package routes

import (
	"your_project/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	protected := r.Group("/")
	protected.Use(middlewares.ApiKeyMiddleware()) // Secure with API key
	{
		protected.GET("/tasks", GetTasks)
		protected.POST("/tasks", CreateTask)
	}
}
