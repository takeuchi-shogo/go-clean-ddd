package entity

import "time"

type UserDetail struct {
	ID        int64     `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	UserID    int64     `gorm:"column:user_id;type:bigint;"`
	FirstName string    `gorm:"column:first_name;type:varchar;size:255;"`
	LastName  string    `gorm:"column:last_name;type:varchar;size:255;"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;"`
}
