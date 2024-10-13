package model

import valueobject "github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/value_object"

type (
	Product struct {
		ID    valueobject.ID
		Title Title
	}
	Title string
)
