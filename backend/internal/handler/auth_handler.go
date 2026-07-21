package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/megatr0n28/autoparts-pro/backend/internal/auth"
)

type AuthHandler struct {
	service *auth.Service
}

func NewAuthHandler(
	service *auth.Service,
) *AuthHandler {

	return &AuthHandler{
		service: service,
	}

}

type LoginRequest struct {
	Email string `json:"email"`

	Password string `json:"password"`
}

func (h *AuthHandler) Login(
	c *gin.Context,
) {

	var request LoginRequest

	if err :=
		c.ShouldBindJSON(
			&request,
		); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid request",
			},
		)

		return
	}

	token, err :=
		h.service.Login(
			c,
			request.Email,
			request.Password,
		)

	if err != nil {

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "invalid credentials",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"token": token,
		},
	)

}
