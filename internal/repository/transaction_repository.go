package repository

import (
	"go-finansia-multi-finance-technical-test/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	Repository[entity.Transaction]
	Log *logrus.Logger
}

func NewTransactionRepository(log *logrus.Logger) *TransactionRepository {
	return &TransactionRepository{
		Log: log,
	}
}

func (r *TransactionRepository) FindByIdAndContactId(tx *gorm.DB, transaction *entity.Transaction, id string, contactId string) error {
	return tx.Where("id = ? AND contact_id = ?", id, contactId).First(transaction).Error
}

func (r *TransactionRepository) FindAllByContactId(tx *gorm.DB, contactId string) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	if err := tx.Where("contact_id = ?", contactId).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) SumOtrByContactAndTenor(tx *gorm.DB, contactId string, tenor int) (int64, error) {
	var total int64
	if err := tx.Model(&entity.Transaction{}).
		Where("contact_id = ? AND tenor = ?", contactId, tenor).
		Select("COALESCE(SUM(otr),0)").
		Scan(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (r *TransactionRepository) SumOtrByContactAndTenorExcludeId(tx *gorm.DB, contactId string, tenor int, excludeId string) (int64, error) {
	var total int64
	if err := tx.Model(&entity.Transaction{}).
		Where("contact_id = ? AND tenor = ? AND id <> ?", contactId, tenor, excludeId).
		Select("COALESCE(SUM(otr),0)").
		Scan(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
