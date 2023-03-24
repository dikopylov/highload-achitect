package handlers

import (
	"errors"
	"github.com/dikopylov/highload-architect/internal/handlers/middleware"
	"github.com/dikopylov/highload-architect/internal/model/users"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserSearchRequest struct {
	FirstName string `form:"first_name" binding:"required"`
	LastName  string `form:"last_name" binding:"required"`
}

type UserSearchResponse struct {
	ID        string     `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name" `
	Biography string     `json:"biography" `
	City      string     `json:"city" `
	Birthdate *time.Time `json:"birthdate"`
	Age       uint       `json:"age" `
}

func (s *implServer) UserSearch(c *gin.Context) {
	var request UserSearchRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, &FailedRequest{
			Message:   err.Error(),
			RequestID: middleware.GetRequestID(c),
			Code:      http.StatusBadRequest,
		})

		return
	}

	userList, err := s.userService.SearchUser(c, &users.SearchUserSpec{
		FirstName: request.FirstName,
		LastName:  request.LastName,
	})
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

	response := make([]*UserSearchResponse, 0, len(userList))

	for _, user := range userList {
		response = append(response, &UserSearchResponse{
			ID:        user.ID.String(),
			Birthdate: user.Birthdate,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Biography: user.Biography,
			City:      user.City,
			Age:       user.Age,
		})
	}

	c.JSON(http.StatusOK, response)
}
