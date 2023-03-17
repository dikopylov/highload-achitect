package handlers

import (
	"github.com/dikopylov/highload-architect/internal/model/users"
	"github.com/gin-gonic/gin"
)

type HTTPServer interface {
	UserRegister(c *gin.Context)
	Login(c *gin.Context)
	GetUserByID(c *gin.Context)
}

type implServer struct {
	userService users.Service
}

func NewHTTPServer(userService users.Service) HTTPServer {
	return &implServer{userService: userService}
}
