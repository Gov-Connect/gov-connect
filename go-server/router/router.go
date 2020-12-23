package router

import (
	"go-server/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router is exported and used in main.go
// func Router() *mux.Router {

// 	router := mux.NewRouter()

// 	// router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
// 	// router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
// 	// router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
// 	// router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
// 	// router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
// 	// router.HandleFunc("/api/deleteAllTask", middleware.DeleteAllTask).Methods("DELETE", "OPTIONS")

// 	return router
// }

// CORSMiddleware handles cors headers for security
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT")

		// if c.Request.Method == "OPTIONS" {
		// 	c.AbortWithStatus(204)
		// 	return
		// }

		c.Next()
	}
}

//NewRouter is exported and used in main.go
func NewRouter() *gin.Engine {
	// new apis
	r := gin.Default()

	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowAllOrigins = true
	// corsConfig.AllowCredentials = true
	// corsConfig.AddAllowedMethods("OPTIONS")
	// corsConfig.AddExposedHeaders("Access-Control-Allow-Origin")
	// r.Use(cors.New(corsConfig))

	r.Use(CORSMiddleware())
	// Setup route group for the API
	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.GET("/localreps", middleware.LocalRepsHandler)
	api.GET("/task", middleware.GetAllTask)
	api.POST("/task", middleware.CreateTask)
	api.PUT("/task/{id}", middleware.TaskComplete)
	api.PUT("/undoTask/{id}", middleware.UndoTask)
	api.OPTIONS("/deleteTask/{id}", middleware.DeleteTask)
	api.OPTIONS("/deleteAllTask", middleware.DeleteAllTask)

	return r
}
