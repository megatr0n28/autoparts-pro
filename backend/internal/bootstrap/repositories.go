package bootstrap

import (
	"gorm.io/gorm"

	"github.com/megatr0n28/autoparts-pro/backend/internal/repository"
	postgresRepo "github.com/megatr0n28/autoparts-pro/backend/internal/repository/postgres"
)

type Repositories struct {
	Transaction   *repository.TransactionManager
	User          repository.UserRepository
	RefreshTokens repository.RefreshTokenRepository
	Customer      repository.CustomerRepository
	Vehicle       repository.VehicleRepository
}

func NewRepositories(
	db *gorm.DB,
) *Repositories {

	return &Repositories{

		Transaction:   repository.NewTransactionManager(db),
		User:          repository.NewUserRepository(db),
		RefreshTokens: postgresRepo.NewRefreshTokenRepository(db),
		Customer:      postgresRepo.NewCustomerRepository(db),
		Vehicle:       postgresRepo.NewVehicleRepository(db),
	}

}
