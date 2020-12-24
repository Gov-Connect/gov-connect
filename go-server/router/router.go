package router

import (
	"go-server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware handles cors headers for security
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")
		c.Header("Context-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

//Router is exported and used in main.go
func Router() *gin.Engine {
	// new apis
	r := gin.Default()

	r.Use(CORSMiddleware())

	// Setup route group for the API
	api := r.Group("/api")

	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// new apis
	api.GET("/local-reps", middleware.LocalRepsHandler)
	api.POST("/local-reps/edit", middleware.EditLocalRep)
	api.GET("/top-reps", middleware.GetTopReps)

	// old apis
	api.GET("/task", middleware.GetAllTask)
	api.POST("/task", middleware.CreateTask)
	api.PUT("/task/:id", middleware.TaskComplete)
	api.PUT("/undoTask/:id", middleware.UndoTask)
	api.OPTIONS("/deleteTask/:id", middleware.DeleteTask)
	api.DELETE("/deleteTask/:id", middleware.DeleteTask)
	api.OPTIONS("/deleteAllTask", middleware.DeleteAllTask)
	api.DELETE("/deleteAllTask", middleware.DeleteAllTask)

	return r
}
