package model

import valueobject "github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/value_object"

type (
	Product struct {
		id           valueobject.ID
		productTitle Title
	}
	Title string

	ProductList []Product
)

func NewProduct(
	id int,
	title string,
) Product {
	return Product{
		id:           valueobject.ID(id),
		productTitle: Title(title),
	}
}

func ReconstructorProduct(
	id int,
	title string,
) Product {
	return Product{
		id:           valueobject.ID(id),
		productTitle: Title(title),
	}
}

func (p Product) ID() valueobject.ID {
	return p.id
}

func (p Product) Title() Title {
	return p.productTitle
}
