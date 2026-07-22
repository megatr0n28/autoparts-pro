package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	vehicleDomain "github.com/megatr0n28/autoparts-pro/backend/internal/domain/vehicle"
	"github.com/megatr0n28/autoparts-pro/backend/internal/dto"
	vehicleService "github.com/megatr0n28/autoparts-pro/backend/internal/service/vehicle"
)

type VehicleHandler struct {
	service *vehicleService.Service
}

func NewVehicleHandler(
	s *vehicleService.Service,
) *VehicleHandler {

	return &VehicleHandler{
		service: s,
	}
}

func (h *VehicleHandler) Create(
	c *gin.Context,
) {

	var request dto.CreateVehicleRequest

	if err := c.ShouldBindJSON(&request); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	customerID :=
		uuid.MustParse(
			c.GetString("customer_id"),
		)

	v :=
		&vehicleDomain.Vehicle{

			CustomerID: customerID,

			VIN: request.VIN,

			Year: request.Year,

			Make: request.Make,

			Model: request.Model,

			Trim: request.Trim,

			Engine: request.Engine,

			Drivetrain: request.Drivetrain,

			Transmission: request.Transmission,

			Mileage: request.Mileage,

			Color: request.Color,

			LicensePlate: request.LicensePlate,

			State: request.State,

			IsPrimary: request.IsPrimary,
		}

	err :=
		h.service.Create(
			c,
			v,
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
		v,
	)
}

func (h *VehicleHandler) List(
	c *gin.Context,
) {

	customerID :=
		uuid.MustParse(
			c.GetString("customer_id"),
		)

	vehicles, err :=
		h.service.List(
			c,
			customerID,
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
		vehicles,
	)
}
