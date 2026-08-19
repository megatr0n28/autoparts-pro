package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/dto"
	dashboardService "github.com/megatr0n28/autoparts-pro/backend/internal/service"
)

type DashboardHandler struct {
	service *dashboardService.DashboardService
}

func NewDashboardHandler(
	service *dashboardService.DashboardService,
) *DashboardHandler {
	return &DashboardHandler{
		service: service,
	}
}

func (h *DashboardHandler) Get(c *gin.Context) {
	userID := c.GetString("user_id")
	customerID, err := uuid.Parse(c.GetString("customer_id"))
	if err != nil || userID == "" {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{"error": "unauthorized"},
		)
		return
	}

	dashboard, err := h.service.Get(c, userID, customerID)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, dashboardService.ErrDashboardAccessDenied) {
			status = http.StatusForbidden
		}

		c.JSON(
			status,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		dto.DashboardResponse{
			Customer: dto.CustomerResponse{
				ID:           dashboard.Customer.ID.String(),
				FirstName:    dashboard.Customer.FirstName,
				LastName:     dashboard.Customer.LastName,
				Phone:        dashboard.Customer.Phone,
				AddressLine1: dashboard.Customer.AddressLine1,
				AddressLine2: dashboard.Customer.AddressLine2,
				City:         dashboard.Customer.City,
				State:        dashboard.Customer.State,
				PostalCode:   dashboard.Customer.PostalCode,
				Country:      dashboard.Customer.Country,
			},
			Vehicles:     dashboard.Vehicles,
			VehicleCount: len(dashboard.Vehicles),
		},
	)
}
