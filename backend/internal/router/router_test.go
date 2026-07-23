package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/megatr0n28/autoparts-pro/backend/internal/testutil"
)

func TestHealthRoute(t *testing.T) {

	engine :=
		testutil.NewTestRouter()

	req :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/v1/health",
			nil,
		)

	resp :=
		httptest.NewRecorder()

	engine.ServeHTTP(
		resp,
		req,
	)

	assert.Equal(
		t,
		http.StatusOK,
		resp.Code,
	)

}

func TestProtectedRoutesRequireJWT(t *testing.T) {

	engine :=
		testutil.NewTestRouter()

	routes :=
		[]string{

			"/api/v1/users/me",

			"/api/v1/customers/me",

			"/api/v1/vehicles",

			"/api/v1/parts/search",
		}

	for _, path := range routes {

		t.Run(
			path,
			func(t *testing.T) {

				req :=
					httptest.NewRequest(
						http.MethodGet,
						path,
						nil,
					)

				resp :=
					httptest.NewRecorder()

				engine.ServeHTTP(
					resp,
					req,
				)

				assert.Equal(
					t,
					http.StatusUnauthorized,
					resp.Code,
				)

			},
		)
	}

}
