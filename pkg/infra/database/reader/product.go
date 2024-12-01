package reader

import (
	"context"

	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/model"
)

type ProductReader struct{}

func NewProductReader() ProductReader {
	return ProductReader{}
}

func (p ProductReader) Get(ctx context.Context, id int) (*model.Product, error) {
	return &model.Product{}, nil
}
