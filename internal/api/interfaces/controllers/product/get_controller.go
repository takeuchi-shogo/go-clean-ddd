package product

import (
	"net/http"
	"strconv"

	"github.com/takeuchi-shogo/go-clean-ddd/internal/api/interfaces/controllers"
	"github.com/takeuchi-shogo/go-clean-ddd/internal/api/usecase/product"
)

type GetProductController struct {
}

func NewGetProductController() GetProductController {
	return GetProductController{}
}

func (pc GetProductController) Get(ctx controllers.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}

	pu := product.NewGetProductUsecase()
	p, err := pu.Get(ctx.Request().Context(), product.GetProductUsecaseInput{
		ID: id,
	})
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, p)
}
