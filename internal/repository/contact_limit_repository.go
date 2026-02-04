package repository

import (
	"go-finansia-multi-finance-technical-test/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ContactLimitRepository struct {
	Repository[entity.ContactLimit]
	Log *logrus.Logger
}

func NewContactLimitRepository(log *logrus.Logger) *ContactLimitRepository {
	return &ContactLimitRepository{
		Log: log,
	}
}

func (r *ContactLimitRepository) FindByIdAndContactId(tx *gorm.DB, limit *entity.ContactLimit, id string, contactId string) error {
	return tx.Where("id = ? AND contact_id = ?", id, contactId).First(limit).Error
}

func (r *ContactLimitRepository) FindByContactIdAndTenor(tx *gorm.DB, limit *entity.ContactLimit, contactId string, tenor int) error {
	return tx.Where("contact_id = ? AND tenor = ?", contactId, tenor).First(limit).Error
}

func (r *ContactLimitRepository) FindAllByContactId(tx *gorm.DB, contactId string) ([]entity.ContactLimit, error) {
	var limits []entity.ContactLimit
	if err := tx.Where("contact_id = ?", contactId).Find(&limits).Error; err != nil {
		return nil, err
	}
	return limits, nil
}
