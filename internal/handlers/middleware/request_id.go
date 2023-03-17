package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDParam = "request_id"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set example variable
		c.Set(RequestIDParam, uuid.NewString())

		// before request

		c.Next()
	}
}

func GetRequestID(c *gin.Context) string {
	value, _ := c.Get(RequestIDParam)

	return value.(string)
}
