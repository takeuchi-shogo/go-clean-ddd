package entity

import "time"

type Product struct {
	ID            int64         `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	ProductDetail ProductDetail `gorm:"foreignKey:ProductID;references:ID"`
	CreatedAt     time.Time     `gorm:"column:created_at;type:datetime;"`
	UpdatedAt     time.Time     `gorm:"column:updated_at;type:datetime;"`
}
