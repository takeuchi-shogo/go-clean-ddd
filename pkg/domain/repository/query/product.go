package query

import (
	"context"

	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/model"
)

type ProductQuery interface {
	Get(context.Context, int) (*model.Product, error)
}
