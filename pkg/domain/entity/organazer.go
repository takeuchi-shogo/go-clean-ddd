package entity

import "time"

type Organazer struct {
	ID              int64           `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;"`
	OrganazerDetail OrganazerDetail `gorm:"foreignKey:OrganazerID;references:ID"`
	CreatedAt       time.Time       `gorm:"column:created_at;type:datetime;"`
	UpdatedAt       time.Time       `gorm:"column:updated_at;type:datetime;"`
}
