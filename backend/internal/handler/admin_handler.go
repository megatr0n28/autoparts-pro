package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/megatr0n28/autoparts-pro/backend/internal/dto"
	adminService "github.com/megatr0n28/autoparts-pro/backend/internal/service"
)

type AdminHandler struct {
	service *adminService.AdminService
}

func NewAdminHandler(service *adminService.AdminService) *AdminHandler {
	return &AdminHandler{service: service}
}

func (h *AdminHandler) Overview(c *gin.Context) {
	overview, err := h.service.Overview(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	users := make([]dto.AdminUserResponse, 0, len(overview.Users))
	for _, item := range overview.Users {
		users = append(users, dto.AdminUserResponse{
			ID: item.ID.String(), FirstName: item.FirstName, LastName: item.LastName,
			Email: item.Email, Role: item.Role, Active: item.Active, CreatedAt: item.CreatedAt,
		})
	}

	customers := make([]dto.AdminCustomerResponse, 0, len(overview.Customers))
	for _, item := range overview.Customers {
		customers = append(customers, dto.AdminCustomerResponse{
			ID: item.ID.String(), UserID: item.UserID.String(), FirstName: item.FirstName,
			LastName: item.LastName, Phone: item.Phone, AddressLine1: item.AddressLine1,
			AddressLine2: item.AddressLine2, City: item.City, State: item.State,
			PostalCode: item.PostalCode, Country: item.Country,
		})
	}

	c.JSON(http.StatusOK, dto.AdminOverviewResponse{
		Users: users, Customers: customers, Vehicles: overview.Vehicles,
		Invoices: []dto.AdminInvoiceResponse{}, InvoiceManagementEnabled: false,
	})
}

func (h *AdminHandler) UpdateUserRole(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var request dto.UpdateUserRoleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role is required"})
		return
	}

	if err := h.service.UpdateUserRole(c, id, request.Role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user role updated"})
}

func (h *AdminHandler) UpdateVehicle(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid vehicle id"})
		return
	}

	var request dto.UpdateVehicleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid vehicle request"})
		return
	}

	if err := h.service.UpdateVehicle(c, id, request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "vehicle updated"})
}

func (h *AdminHandler) DeleteVehicle(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid vehicle id"})
		return
	}

	if err := h.service.DeleteVehicle(c, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *AdminHandler) UpdateCustomer(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer id"})
		return
	}

	var request dto.UpdateCustomerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid customer request"})
		return
	}

	if err := h.service.UpdateCustomer(c, id, request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer updated"})
}
