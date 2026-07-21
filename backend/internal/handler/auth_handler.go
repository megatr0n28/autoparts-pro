package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/megatr0n28/autoparts-pro/backend/internal/auth"
	domainUser "github.com/megatr0n28/autoparts-pro/backend/internal/domain/user"
	"github.com/megatr0n28/autoparts-pro/backend/internal/dto"
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

	accessToken,
		refreshToken,
		err :=
		h.service.Login(
			c,
			request.Email,
			request.Password,
			c.GetHeader("User-Agent"),
			c.ClientIP(),
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
		dto.LoginResponse{

			AccessToken: accessToken,

			RefreshToken: refreshToken,
		},
	)

}

func (h *AuthHandler) Register(
	c *gin.Context,
) {

	var request dto.RegisterRequest

	if err :=
		c.ShouldBindJSON(
			&request,
		); err != nil {

		c.JSON(
			400,
			gin.H{
				"error": "invalid request",
			},
		)

		return
	}

	u :=
		&domainUser.User{

			FirstName: request.FirstName,

			LastName: request.LastName,

			Email: request.Email,
		}

	err :=
		h.service.Register(
			c,
			u,
			request.Password,
		)

	if err != nil {

		c.JSON(
			500,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		201,
		gin.H{
			"message": "user created",
		},
	)

}

func (h *AuthHandler) Refresh(
	c *gin.Context,
) {

	var request dto.RefreshRequest

	if err :=
		c.ShouldBindJSON(&request); err != nil {

		c.JSON(
			400,
			gin.H{
				"error": "invalid request",
			},
		)

		return
	}

	accessToken,
		refreshToken,
		err :=
		h.service.Refresh(
			c,
			request.RefreshToken,
			c.GetHeader("User-Agent"),
			c.ClientIP(),
		)

	if err != nil {

		c.JSON(
			401,
			gin.H{
				"error": "invalid refresh token",
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		dto.TokenResponse{

			AccessToken: accessToken,

			RefreshToken: refreshToken,
		},
	)

}

func (h *AuthHandler) Logout(
	c *gin.Context,
) {

	var request dto.RefreshRequest

	if err :=
		c.ShouldBindJSON(&request); err != nil {

		c.JSON(
			400,
			gin.H{
				"error": "invalid request",
			},
		)

		return
	}

	err :=
		h.service.Logout(
			c,
			request.RefreshToken,
		)

	if err != nil {

		c.JSON(
			500,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		200,
		gin.H{
			"message": "logged out",
		},
	)

}

func (h *AuthHandler) LogoutAll(
	c *gin.Context,
) {

	userID,
		exists :=
		c.Get("user_id")

	if !exists {

		c.JSON(
			401,
			gin.H{
				"error": "unauthorized",
			},
		)

		return
	}

	err :=
		h.service.LogoutAll(
			c,
			uuid.MustParse(
				userID.(string),
			),
		)

	if err != nil {

		c.JSON(
			500,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		200,
		gin.H{
			"message": "all sessions terminated",
		},
	)

}
