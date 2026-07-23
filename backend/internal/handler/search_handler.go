package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/service/search"
)

type SearchHandler struct {
	service *search.Service
}

func NewSearchHandler(
	service *search.Service,
) *SearchHandler {

	return &SearchHandler{
		service: service,
	}
}

// SearchParts
//
// GET /api/v1/search/parts
//
// Query params:
// vehicle_id
// query
func (h *SearchHandler) SearchParts(
	c *gin.Context,
) {

	vehicleIDParam :=
		c.Query(
			"vehicle_id",
		)

	query :=
		c.Query(
			"query",
		)

	if vehicleIDParam == "" {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "vehicle_id required",
			},
		)

		return
	}

	vehicleID, err :=
		uuid.Parse(
			vehicleIDParam,
		)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid vehicle id",
			},
		)

		return
	}

	if query == "" {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "search query required",
			},
		)

		return
	}

	results, err :=
		h.service.Search(
			c.Request.Context(),
			vehicleID,
			query,
		)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"results": results,
		},
	)

}
