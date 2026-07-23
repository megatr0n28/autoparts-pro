package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/handler"
	"github.com/megatr0n28/autoparts-pro/backend/internal/provider/mock"
	"github.com/megatr0n28/autoparts-pro/backend/internal/service/search"
	"go.uber.org/zap"
)

func setupRouter() *gin.Engine {

	gin.SetMode(
		gin.TestMode,
	)

	logger :=
		zap.NewNop()

	provider :=
		mock.New()

	searchService :=
		search.New(
			logger,
			provider,
		)

	searchHandler :=
		handler.NewSearchHandler(
			searchService,
		)

	router :=
		gin.New()

	router.GET(
		"/api/v1/parts/search",
		searchHandler.SearchParts,
	)

	return router
}

func TestSearchParts_SmokeTest(t *testing.T) {

	router :=
		setupRouter()

	vehicleID :=
		uuid.New()

	request :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/v1/parts/search?vehicle_id="+vehicleID.String()+"&query=oil%20filter",
			nil,
		)

	response :=
		httptest.NewRecorder()

	router.ServeHTTP(
		response,
		request,
	)

	assert.Equal(
		t,
		http.StatusOK,
		response.Code,
	)

	body :=
		response.Body.String()

	assert.Contains(
		t,
		body,
		"FRAM",
	)

	assert.Contains(
		t,
		body,
		"AutoZone",
	)

}

func TestSearchParts_MissingVehicleID(t *testing.T) {

	router :=
		setupRouter()

	request :=
		httptest.NewRequest(
			http.MethodGet,
			"/api/v1/parts/search?query=oil%20filter",
			nil,
		)

	response :=
		httptest.NewRecorder()

	router.ServeHTTP(
		response,
		request,
	)

	assert.Equal(
		t,
		http.StatusBadRequest,
		response.Code,
	)

}
