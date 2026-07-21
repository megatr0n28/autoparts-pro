package bootstrap

import (
	"gorm.io/gorm"

	"github.com/megatr0n28/autoparts-pro/backend/internal/repository"
)

type Repositories struct {
	Transaction *repository.TransactionManager
	User        repository.UserRepository
}

func NewRepositories(
	db *gorm.DB,
) *Repositories {

	return &Repositories{

		Transaction: repository.NewTransactionManager(db),
		User:        repository.NewUserRepository(db),
	}

}
