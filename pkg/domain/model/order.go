package model

import valueobject "github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/value_object"

type (
	Order struct {
		ID      valueobject.ID
		UserID  valueobject.ID
		Product Product
	}
	OrderList []Order
)
