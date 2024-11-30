package entity

import (
	"time"

	"github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/model"
	valueobject "github.com/takeuchi-shogo/go-clean-ddd/pkg/domain/value_object"
)

type Order struct {
	ID          int64       `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	UserOrder   UserOrder   `gorm:"many2many:user_orders;"`
	OrderDetail OrderDetail `gorm:"foreignKey:OrderID;references:ID"`
	Product     Product     `gorm:"many2many:product_oders"`
	CreatedAt   time.Time   `gorm:"column:created_at;type:datetime;"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;type:datetime;"`
}

func (o Order) ToModel() *model.Order {
	return &model.Order{
		ID:     valueobject.ID(o.ID),
		UserID: valueobject.ID(o.UserOrder.UserID),
	}
}
