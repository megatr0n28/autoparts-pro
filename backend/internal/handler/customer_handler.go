package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/dto"

	customerService "github.com/megatr0n28/autoparts-pro/backend/internal/service"
)

type CustomerHandler struct {
	service *customerService.CustomerService
}

func NewCustomerHandler(
	service *customerService.CustomerService,
) *CustomerHandler {

	return &CustomerHandler{
		service: service,
	}

}

func (h *CustomerHandler) Me(
	c *gin.Context,
) {

	userIDValue, exists :=
		c.Get(
			"user_id",
		)

	if !exists {

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "unauthorized",
			},
		)

		return

	}

	userID :=
		uuid.MustParse(
			userIDValue.(string),
		)

	customer,
		err :=
		h.service.Get(
			c,
			userID,
		)

	if err != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "customer profile not found",
			},
		)

		return

	}

	c.JSON(
		http.StatusOK,
		dto.CustomerResponse{

			ID: customer.ID.String(),

			FirstName: customer.FirstName,

			LastName: customer.LastName,

			Phone: customer.Phone,

			AddressLine1: customer.AddressLine1,

			AddressLine2: customer.AddressLine2,

			City: customer.City,

			State: customer.State,

			PostalCode: customer.PostalCode,

			Country: customer.Country,
		},
	)

}

func (h *CustomerHandler) Update(
	c *gin.Context,
) {

	userIDValue, _ :=
		c.Get(
			"user_id",
		)

	userID :=
		uuid.MustParse(
			userIDValue.(string),
		)

	customer,
		err :=
		h.service.Get(
			c,
			userID,
		)

	if err != nil {

		c.JSON(
			404,
			gin.H{
				"error": "customer profile not found",
			},
		)

		return

	}

	var request dto.UpdateCustomerRequest

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

	customer.FirstName =
		request.FirstName

	customer.LastName =
		request.LastName

	customer.Phone =
		request.Phone

	customer.AddressLine1 =
		request.AddressLine1

	customer.AddressLine2 =
		request.AddressLine2

	customer.City =
		request.City

	customer.State =
		request.State

	customer.PostalCode =
		request.PostalCode

	err =
		h.service.Update(
			c,
			customer,
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
			"message": "profile updated",
		},
	)

}
