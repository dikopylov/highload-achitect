package handlers

import (
	"github.com/dikopylov/highload-architect/internal/handlers/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type FailedRequest struct {
	Message   string `form:"message"`
	RequestID string `form:"request_id"`
	Code      int    `form:"code"`
}

func InitRouter(server HTTPServer) *gin.Engine {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	r.Use(middleware.RequestID())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", server.Login)

	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/register", server.UserRegister)
		userRoutes.GET("/get/:id", server.GetUserByID)
		userRoutes.GET("/search", server.UserSearch)
	}

	r.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}
