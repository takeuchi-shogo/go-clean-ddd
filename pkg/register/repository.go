package register

import (
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/repository"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/infra/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (*Repository, error) {
	return &Repository{
		db: db,
	}, nil
}

func (repo *Repository) NewUserRepository() repository.UserRepository {
	return database.NewUserRepository(repo.db)
}
