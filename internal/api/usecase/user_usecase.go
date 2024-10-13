package usecase

import (
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/model"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/repository"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/register"
)

type GetUserInput struct {
	UserID int
}

type UsecaseOutput struct {
}

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase() *UserUsecase {
	repo := register.New().Repository()
	return &UserUsecase{
		userRepo: repo.NewUserRepository(),
	}
}

func (u *UserUsecase) Get(ipt GetUserInput) (*model.User, error) {
	user, err := u.userRepo.FindByID(ipt.UserID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
