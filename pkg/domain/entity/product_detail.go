package entity

import "time"

type ProductDetail struct {
	ID          int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	ProductID   int64     `gorm:"column:product_id;type:bigint;"`
	Title       string    `gorm:"column:title;type:varchar;size:255;"`
	Desctiption string    `gorm:"column:desctiption;type:text;"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;"`
}
