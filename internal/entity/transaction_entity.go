package entity

import "time"

type Transaction struct {
	ID                string    `gorm:"column:id;primaryKey"`
	ContactId         string    `gorm:"column:contact_id"`
	ContractNumber    string    `gorm:"column:contract_number"`
	Tenor             int       `gorm:"column:tenor"`
	Channel           string    `gorm:"column:channel"`
	Otr               int64     `gorm:"column:otr"`
	AdminFee          int64     `gorm:"column:admin_fee"`
	InstallmentAmount int64     `gorm:"column:installment_amount"`
	InterestAmount    int64     `gorm:"column:interest_amount"`
	AssetName         string    `gorm:"column:asset_name"`
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Contact           Contact   `gorm:"foreignKey:contact_id;references:id"`
}

func (t *Transaction) TableName() string {
	return "transactions"
}
