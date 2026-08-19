package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/megatr0n28/autoparts-pro/backend/internal/middleware"
)

func TestRequireWriteAccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	for _, test := range []struct {
		name string
		role string
		code int
	}{
		{name: "viewer is read only", role: "viewer", code: http.StatusForbidden},
		{name: "editor can write", role: "editor", code: http.StatusOK},
		{name: "admin can write", role: "admin", code: http.StatusOK},
		{name: "legacy user can write", role: "user", code: http.StatusOK},
	} {
		t.Run(test.name, func(t *testing.T) {
			router := gin.New()
			router.Use(func(c *gin.Context) {
				c.Set("role", test.role)
				c.Next()
			})
			router.PUT("/resource", middleware.RequireWriteAccess(), func(c *gin.Context) {
				c.Status(http.StatusOK)
			})

			response := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPut, "/resource", nil)
			router.ServeHTTP(response, request)

			assert.Equal(t, test.code, response.Code)
		})
	}
}
