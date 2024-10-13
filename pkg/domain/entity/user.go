package entity

import (
	"time"

	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/model"
	valueobject "github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/value_object"
)

type User struct {
	ID          int64       `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	UserDetail  UserDetail  `gorm:"foreignKey:UserID;references:ID"`
	UserAccount UserAccount `gorm:"foreignKey:UserID;references:ID"`
	Orders      []Order     `gorm:"many2many:user_orders;"`
	CreatedAt   time.Time   `gorm:"column:created_at;type:datetime;"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;type:datetime;"`
}

func (u User) ToModel() *model.User {
	return &model.User{
		ID:        valueobject.ID(u.ID),
		FirstName: valueobject.FirstName(u.UserDetail.FirstName),
		LastName:  valueobject.LastName(u.UserDetail.LastName),
		Email:     valueobject.Email(u.UserAccount.Email),
	}
}
