package register

import (
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/repository"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/repository/query"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/infra/database"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/infra/database/reader"
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

func (repo *Repository) NewProductQuery() query.ProductQuery {
	return reader.NewProductReader()
}
