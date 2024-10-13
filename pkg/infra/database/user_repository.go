package database

import (
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/entity"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Find() {}

func (u *UserRepository) FindByID(id int) (*model.User, error) {
	user := &entity.User{}
	if err := u.defaultScope().
		Where("id = ?", id).
		First(user).
		Error; err != nil {
		return nil, err
	}
	return user.ToModel(), nil
}

func (u *UserRepository) defaultScope() *gorm.DB {
	return u.db.
		Preload("UserDetail").
		Preload("UserAccount")
}
