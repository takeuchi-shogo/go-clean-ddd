package valueobject

import "errors"

type Email string

func (e Email) IsValid() error {
	if e == Email("") {
		return errors.New("empty email")
	}
	return nil
}
