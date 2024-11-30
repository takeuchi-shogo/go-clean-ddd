package model

import valueobject "github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/value_object"

type (
	Organazer struct {
		ID       valueobject.ID
		Name     Name
		Products ProductList
	}
	Name string
)
