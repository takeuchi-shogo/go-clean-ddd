package model

import valueobject "github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/value_object"

type (
	User struct {
		ID        valueobject.ID
		FirstName valueobject.FirstName
		LastName  valueobject.LastName
		Email     valueobject.Email
		Orders    OrderList
	}
)
