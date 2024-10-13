package controllers

import (
	"net/http"

	"github.com/takeuchi-shogo/go-clean-ddd/internal/api/domain"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/model"
	valueobject "github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/value_object"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) Get(ctx Context) error {
	return ctx.JSON(http.StatusOK, &domain.User{
		User: model.User{
			ID: valueobject.ID(1000),
		},
	})
}
