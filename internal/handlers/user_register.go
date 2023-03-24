package handlers

import (
	"github.com/dikopylov/highload-architect/internal/handlers/middleware"
	"github.com/dikopylov/highload-architect/internal/model/users"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserRegisterRequest struct {
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Biography string    `json:"biography" binding:"required"`
	City      string    `json:"city" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Age       int       `json:"age" binding:"required,gt=0"`
	Birthdate time.Time `json:"birthdate" time_format:"2006-01-02" time_utc:"1" binding:"required"`
}

type UserRegisterResponse struct {
	UserID string `json:"user_id"`
}

func (s *implServer) UserRegister(c *gin.Context) {
	var request UserRegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, &FailedRequest{
			Message:   err.Error(),
			RequestID: middleware.GetRequestID(c),
			Code:      http.StatusBadRequest,
		})

		return
	}

	user := &users.User{
		Birthdate: &request.Birthdate,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Biography: request.Biography,
		City:      request.City,
		Password:  request.Password,
		Age:       uint(request.Age),
	}

	err := s.userService.Register(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &FailedRequest{
			Message:   err.Error(),
			RequestID: middleware.GetRequestID(c),
			Code:      http.StatusInternalServerError,
		})

		return
	}

	c.JSON(http.StatusOK, UserRegisterResponse{UserID: user.ID.String()})
}
