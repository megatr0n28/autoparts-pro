package testutil

import (
	"github.com/megatr0n28/autoparts-pro/backend/internal/handler"
)

func NewUserHandler() *handler.UserHandler {

	return handler.NewUserHandler()

}

func NewAuthHandler() *handler.AuthHandler {

	return handler.NewTestAuthHandler()

}

func NewCustomerHandler() *handler.CustomerHandler {

	return handler.NewTestCustomerHandler()

}

func NewVehicleHandler() *handler.VehicleHandler {

	return handler.NewTestVehicleHandler()

}

func NewSearchHandler() *handler.SearchHandler {

	return handler.NewTestSearchHandler()

}
