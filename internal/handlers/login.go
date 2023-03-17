package handlers

import (
	"github.com/dikopylov/highload-architect/internal/handlers/middleware"
	"github.com/dikopylov/highload-architect/internal/model/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type UserLoginRequest struct {
	ID       string `json:"id" binding:"required,uuid"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

func (s *implServer) Login(c *gin.Context) {
	var request UserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
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
	token, err := s.userService.Login(c, types.MakeUserIDByUUID(id), request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &FailedRequest{
			Message:   err.Error(),
			RequestID: middleware.GetRequestID(c),
			Code:      http.StatusInternalServerError,
		})

		return
	}

	c.JSON(http.StatusOK, &UserLoginResponse{
		Token: token.String(),
	})
}
