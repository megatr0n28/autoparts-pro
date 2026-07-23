package testutil

import (
	"github.com/gin-gonic/gin"

	"github.com/megatr0n28/autoparts-pro/backend/internal/auth"
	"github.com/megatr0n28/autoparts-pro/backend/internal/router"
)

func NewTestRouter() *gin.Engine {

	jwtManager :=
		auth.NewJWTManager(
			"test-secret",
			3600,
		)

	userHandler :=
		NewUserHandler()

	authHandler :=
		NewAuthHandler()

	customerHandler :=
		NewCustomerHandler()

	vehicleHandler :=
		NewVehicleHandler()

	customerRepository :=
		NewCustomerRepository()

	searchHandler :=
		NewSearchHandler()

	return router.New(
		jwtManager,
		userHandler,
		authHandler,
		customerHandler,
		vehicleHandler,
		customerRepository,
		searchHandler,
	)

}
