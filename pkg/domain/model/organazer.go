package model

import valueobject "github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/value_object"

type (
	Organazer struct {
		id       valueobject.ID
		name     Name
		products ProductList
	}
	Name string
)

func NewOrganazer(
	id int,
	name string,
	products []Product,
) Organazer {
	return Organazer{
		id:       valueobject.ID(id),
		name:     Name(name),
		products: ProductList(products),
	}
}

func (o Organazer) ID() valueobject.ID {
	return o.id
}

func (o Organazer) Name() Name {
	return o.name
}
