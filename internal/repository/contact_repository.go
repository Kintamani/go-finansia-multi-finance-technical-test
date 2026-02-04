package repository

import (
	"go-finansia-multi-finance-technical-test/internal/entity"
	"go-finansia-multi-finance-technical-test/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ContactRepository struct {
	Repository[entity.Contact]
	Log *logrus.Logger
}

func NewContactRepository(log *logrus.Logger) *ContactRepository {
	return &ContactRepository{
		Log: log,
	}
}

func (r *ContactRepository) FindByIdAndUserId(db *gorm.DB, contact *entity.Contact, id string, userId int64) error {
	return db.Where("id = ? AND user_id = ?", id, userId).Take(contact).Error
}

func (r *ContactRepository) Search(db *gorm.DB, request *model.SearchContactRequest) ([]entity.Contact, int64, error) {
	var contacts []entity.Contact
	if err := db.Scopes(r.FilterContact(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&contacts).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Contact{}).Scopes(r.FilterContact(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return contacts, total, nil
}

func (r *ContactRepository) FilterContact(request *model.SearchContactRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("user_id = ?", request.UserId)

		if nik := request.NIK; nik != "" {
			tx = tx.Where("nik = ?", nik)
		}

		if fullName := request.FullName; fullName != "" {
			fullName = "%" + fullName + "%"
			tx = tx.Where("full_name LIKE ?", fullName)
		}

		if legalName := request.LegalName; legalName != "" {
			legalName = "%" + legalName + "%"
			tx = tx.Where("legal_name LIKE ?", legalName)
		}

		return tx
	}
}
