package command

import (
	"context"

	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/model"
)

type ProductCommand interface {
	Create(context.Context, *model.Product) error
}
