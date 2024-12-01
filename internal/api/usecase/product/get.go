package product

import "context"

type GetProductUsecaseInput struct {
	ID int
}

type GetProductUsecaseOutput struct{}

type GetProductUsecase struct {
}

func NewGetProductUsecase() GetProductUsecase {
	return GetProductUsecase{}
}

func (pu GetProductUsecase) Get(
	ctx context.Context,
	ipt GetProductUsecaseInput,
) (GetProductUsecaseOutput, error) {
	return GetProductUsecaseOutput{}, nil
}
