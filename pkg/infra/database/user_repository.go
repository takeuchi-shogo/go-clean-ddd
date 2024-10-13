package database

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Find() {}

func (u *UserRepository) FindByID() {}
