package controllers

import (
	"net/http"
	"strconv"

	"github.com/takeuchi-shogo/go-clean-ddd/internal/api/domain"
	"github.com/takeuchi-shogo/go-clean-ddd/internal/api/usecase"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/model"
	valueobject "github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/value_object"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) GetList(ctx Context) error {
	return ctx.JSON(http.StatusOK, &domain.User{
		User: &model.User{
			ID: valueobject.ID(1000),
		},
	})
}

func (u *UserController) Get(ctx Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	uc := usecase.NewUserUsecase()
	user, res := uc.Get(usecase.GetUserInput{
		UserID: id,
	})
	if res != nil {
		return ctx.JSON(http.StatusNotFound, nil)
	}
	return ctx.JSON(http.StatusOK, user)
}
