package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/dto"
	searchService "github.com/megatr0n28/autoparts-pro/backend/internal/service/search"
)

type SearchHandler struct {
	service *searchService.Service
}

func NewSearchHandler(
	s *searchService.Service,
) *SearchHandler {

	return &SearchHandler{
		service: s,
	}
}

func (h *SearchHandler) Search(
	c *gin.Context,
) {

	var request dto.PartSearchRequest

	if err :=
		c.ShouldBindQuery(
			&request,
		); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	vehicleID, err :=
		uuid.Parse(
			request.VehicleID,
		)

	if err != nil {

		c.JSON(
			400,
			gin.H{
				"error": "invalid vehicle id",
			},
		)

		return
	}

	results, err :=
		h.service.Search(
			c,
			vehicleID,
			request.Query,
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
		results,
	)
}
