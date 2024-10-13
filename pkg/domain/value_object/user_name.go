package valueobject

import "errors"

type (
	FirstName string
	LastName  string
)

var (
	ErrEmptyName = errors.New("empty name")
)

func (u FirstName) CheckEmpty() error {
	if u == FirstName("") {
		return ErrEmptyName
	}
	return nil
}

func (u LastName) CheckEmpty() error {
	if u == LastName("") {
		return ErrEmptyName
	}
	return nil
}
