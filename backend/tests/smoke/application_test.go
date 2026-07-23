package smoke_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/megatr0n28/autoparts-pro/backend/internal/testutil"
)

func TestApplicationHealth(t *testing.T) {

	engine :=
		testutil.NewTestRouter()

	request :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/v1/health",
			nil,
		)

	response :=
		httptest.NewRecorder()

	engine.ServeHTTP(
		response,
		request,
	)

	assert.Equal(
		t,
		http.StatusOK,
		response.Code,
	)

}

func TestApplicationAuthRoutesExist(t *testing.T) {

	engine :=
		testutil.NewTestRouter()

	tests :=
		[]struct {
			name   string
			method string
			path   string
		}{
			{
				name:   "register",
				method: http.MethodPost,
				path:   "/api/v1/auth/register",
			},
			{
				name:   "login",
				method: http.MethodPost,
				path:   "/api/v1/auth/login",
			},
			{
				name:   "refresh",
				method: http.MethodPost,
				path:   "/api/v1/auth/refresh",
			},
		}

	for _, test := range tests {

		t.Run(
			test.name,
			func(t *testing.T) {

				request :=
					httptest.NewRequest(
						test.method,
						test.path,
						nil,
					)

				response :=
					httptest.NewRecorder()

				engine.ServeHTTP(
					response,
					request,
				)

				// route exists
				assert.NotEqual(
					t,
					http.StatusNotFound,
					response.Code,
				)

			},
		)
	}

}

func TestProtectedRoutesRequireAuthentication(t *testing.T) {

	engine :=
		testutil.NewTestRouter()

	tests :=
		[]string{

			"/api/v1/users/me",

			"/api/v1/customers/me",

			"/api/v1/vehicles",

			"/api/v1/parts/search",
		}

	for _, path := range tests {

		t.Run(
			path,
			func(t *testing.T) {

				request :=
					httptest.NewRequest(
						http.MethodGet,
						path,
						nil,
					)

				response :=
					httptest.NewRecorder()

				engine.ServeHTTP(
					response,
					request,
				)

				assert.Equal(
					t,
					http.StatusUnauthorized,
					response.Code,
				)

			},
		)

	}

}
