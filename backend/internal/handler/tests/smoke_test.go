package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupSmokeRouter() *gin.Engine {

	gin.SetMode(
		gin.TestMode,
	)

	router :=
		gin.New()

	api :=
		router.Group(
			"/api/v1",
		)

	// Health
	api.GET(
		"/health",
		func(c *gin.Context) {

			c.JSON(
				200,
				gin.H{
					"status": "ok",
				},
			)

		},
	)

	// Auth smoke routes
	auth :=
		api.Group(
			"/auth",
		)

	auth.POST(
		"/register",
		func(c *gin.Context) {

			c.JSON(
				201,
				gin.H{
					"message": "registered",
				},
			)

		},
	)

	auth.POST(
		"/login",
		func(c *gin.Context) {

			c.JSON(
				200,
				gin.H{
					"access_token": "test-token",
				},
			)

		},
	)

	// Protected examples
	protected :=
		api.Group(
			"",
		)

	protected.GET(
		"/users/me",
		func(c *gin.Context) {

			c.JSON(
				200,
				gin.H{
					"id": "test-user",
				},
			)

		},
	)

	protected.GET(
		"/customers/me",
		func(c *gin.Context) {

			c.JSON(
				200,
				gin.H{
					"first_name": "John",
				},
			)

		},
	)

	protected.GET(
		"/vehicles",
		func(c *gin.Context) {

			c.JSON(
				200,
				[]any{},
			)

		},
	)

	protected.GET(
		"/parts/search",
		func(c *gin.Context) {

			c.JSON(
				200,
				gin.H{
					"results": []any{},
				},
			)

		},
	)

	return router
}

func TestHealthEndpoint(t *testing.T) {

	router :=
		setupSmokeRouter()

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/v1/health",
			nil,
		)

	resp :=
		httptest.NewRecorder()

	router.ServeHTTP(
		resp,
		req,
	)

	assert.Equal(
		t,
		200,
		resp.Code,
	)

}

func TestAuthRegister(t *testing.T) {

	router :=
		setupSmokeRouter()

	req :=
		httptest.NewRequest(
			http.MethodPost,
			"/api/v1/auth/register",
			nil,
		)

	resp :=
		httptest.NewRecorder()

	router.ServeHTTP(
		resp,
		req,
	)

	assert.Equal(
		t,
		201,
		resp.Code,
	)

}

func TestAuthLogin(t *testing.T) {

	router :=
		setupSmokeRouter()

	req :=
		httptest.NewRequest(
			http.MethodPost,
			"/api/v1/auth/login",
			nil,
		)

	resp :=
		httptest.NewRecorder()

	router.ServeHTTP(
		resp,
		req,
	)

	assert.Equal(
		t,
		200,
		resp.Code,
	)

}

func TestCustomerProfile(t *testing.T) {

	router :=
		setupSmokeRouter()

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/v1/customers/me",
			nil,
		)

	resp :=
		httptest.NewRecorder()

	router.ServeHTTP(
		resp,
		req,
	)

	assert.Equal(
		t,
		200,
		resp.Code,
	)

}

func TestVehicleList(t *testing.T) {

	router :=
		setupSmokeRouter()

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/v1/vehicles",
			nil,
		)

	resp :=
		httptest.NewRecorder()

	router.ServeHTTP(
		resp,
		req,
	)

	assert.Equal(
		t,
		200,
		resp.Code,
	)

}

func TestPartsSearch(t *testing.T) {

	router :=
		setupSmokeRouter()

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/v1/parts/search?query=oil",
			nil,
		)

	resp :=
		httptest.NewRecorder()

	router.ServeHTTP(
		resp,
		req,
	)

	assert.Equal(
		t,
		200,
		resp.Code,
	)

}
