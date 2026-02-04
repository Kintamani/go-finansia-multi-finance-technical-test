package entity

import "time"

type ContactLimit struct {
	ID          string    `gorm:"column:id;primaryKey"`
	ContactId   string    `gorm:"column:contact_id"`
	Tenor       int       `gorm:"column:tenor"`
	LimitAmount int64     `gorm:"column:limit_amount"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Contact     Contact   `gorm:"foreignKey:contact_id;references:id"`
}

func (c *ContactLimit) TableName() string {
	return "contact_limits"
}
