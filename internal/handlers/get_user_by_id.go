package handlers

import (
	"errors"
	"github.com/dikopylov/highload-architect/internal/handlers/middleware"
	"github.com/dikopylov/highload-architect/internal/model/types"
	"github.com/dikopylov/highload-architect/internal/model/users"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type GetUserByIDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type GetUserByIDResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name" `
	Biography string    `json:"biography" `
	City      string    `json:"city" `
	Birthdate time.Time `json:"birthdate"`
	Age       uint      `json:"age" `
}

func (s *implServer) GetUserByID(c *gin.Context) {
	var request GetUserByIDRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(http.StatusBadRequest, &FailedRequest{
			Message:   err.Error(),
			RequestID: middleware.GetRequestID(c),
			Code:      http.StatusBadRequest,
		})

		return
	}

	id, errUUID := uuid.Parse(request.ID)
	if errUUID != nil {
		c.JSON(http.StatusBadRequest, &FailedRequest{
			Message:   errUUID.Error(),
			RequestID: middleware.GetRequestID(c),
			Code:      http.StatusBadRequest,
		})

		return
	}

	user, err := s.userService.GetUserByID(c, types.MakeUserIDByUUID(id))
	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, &FailedRequest{
				Message:   err.Error(),
				RequestID: middleware.GetRequestID(c),
				Code:      http.StatusNotFound,
			})

			return
		}

		c.JSON(http.StatusInternalServerError, &FailedRequest{
			Message:   err.Error(),
			RequestID: middleware.GetRequestID(c),
			Code:      http.StatusInternalServerError,
		})

		return
	}

	c.JSON(http.StatusOK, &GetUserByIDResponse{
		ID:        user.ID.String(),
		Birthdate: *user.Birthdate,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Biography: user.Biography,
		City:      user.City,
		Age:       user.Age,
	})
}
