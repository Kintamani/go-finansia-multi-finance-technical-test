package entity

import "time"

type Contact struct {
	ID          string    `gorm:"column:id;primaryKey"`
	NIK         string    `gorm:"column:nik"`
	FullName    string    `gorm:"column:full_name"`
	LegalName   string    `gorm:"column:legal_name"`
	BirthPlace  string    `gorm:"column:birth_place"`
	BirthDate   string    `gorm:"column:birth_date"`
	Salary      int64     `gorm:"column:salary"`
	KtpPhoto    string    `gorm:"column:ktp_photo"`
	SelfiePhoto string    `gorm:"column:selfie_photo"`
	UserId      int64     `gorm:"column:user_id"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	User        User      `gorm:"foreignKey:user_id;references:id"`
	Addresses   []Address `gorm:"foreignKey:contact_id;references:id"`
}

func (c *Contact) TableName() string {
	return "contacts"
}
