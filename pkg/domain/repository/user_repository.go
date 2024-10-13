package repository

import "github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/model"

type UserRepository interface {
	Find()
	FindByID(id int) (*model.User, error)
}
