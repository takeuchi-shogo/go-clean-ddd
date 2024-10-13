package entity

import "time"

type OrderDetail struct {
	ID          int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	OrderID     int64     `gorm:"column:order_id;type:bigint;"`
	OrganazerID int64     `gorm:"column:organazer_id;type:bigint;"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;"`
}
