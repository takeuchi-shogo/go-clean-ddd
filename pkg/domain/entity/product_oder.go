package entity

import "time"

type ProductOrder struct {
	ID        int64
	ProductID int64     `gorm:"column:product_id;type:bigint;"`
	OrderID   int64     `gorm:"column:order_id;type:bigint;"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;"`
}
