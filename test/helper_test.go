package test

import (
	"fmt"
	"go-finansia-multi-finance-technical-test/internal/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func ClearAll() {
	ClearTransactions()
	ClearContactLimits()
	ClearAddresses()
	ClearContact()
	ClearUsers()
}

func ClearUsers() {
	err := db.Where("id is not null").Delete(&entity.User{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}

func ClearContact() {
	err := db.Where("id is not null").Delete(&entity.Contact{}).Error
	if err != nil {
		log.Fatalf("Failed clear contact data : %+v", err)
	}
}

func ClearContactLimits() {
	err := db.Where("id is not null").Delete(&entity.ContactLimit{}).Error
	if err != nil {
		log.Fatalf("Failed clear contact limit data : %+v", err)
	}
}

func ClearTransactions() {
	err := db.Where("id is not null").Delete(&entity.Transaction{}).Error
	if err != nil {
		log.Fatalf("Failed clear transaction data : %+v", err)
	}
}

func ClearAddresses() {
	err := db.Where("id is not null").Delete(&entity.Address{}).Error
	if err != nil {
		log.Fatalf("Failed clear address data : %+v", err)
	}
}

func CreateContacts(user *entity.User, total int) {
	for i := 0; i < total; i++ {
		contact := &entity.Contact{
			ID:          uuid.NewString(),
			NIK:         fmt.Sprintf("320000000000%04d", i),
			FullName:    fmt.Sprintf("Contact %d", i),
			LegalName:   fmt.Sprintf("Contact %d", i),
			BirthPlace:  "Jakarta",
			BirthDate:   "1990-01-01",
			Salary:      5000000 + int64(i),
			KtpPhoto:    fmt.Sprintf("https://example.com/ktp/%d.jpg", i),
			SelfiePhoto: fmt.Sprintf("https://example.com/selfie/%d.jpg", i),
			UserId:      user.ID,
		}
		err := db.Create(contact).Error
		if err != nil {
			log.Fatalf("Failed create contact data : %+v", err)
		}
	}
}

func CreateContactLimit(t *testing.T, contact *entity.Contact, tenor int, limitAmount int64) *entity.ContactLimit {
	limit := &entity.ContactLimit{
		ID:          uuid.NewString(),
		ContactId:   contact.ID,
		Tenor:       tenor,
		LimitAmount: limitAmount,
	}
	err := db.Create(limit).Error
	assert.Nil(t, err)
	return limit
}

func CreateTransactions(t *testing.T, contact *entity.Contact, total int, tenor int) {
	for i := 0; i < total; i++ {
		transaction := &entity.Transaction{
			ID:                uuid.NewString(),
			ContactId:         contact.ID,
			ContractNumber:    fmt.Sprintf("CN-%d-%d", tenor, i),
			Tenor:             tenor,
			Channel:           "ecommerce",
			Otr:               100000 + int64(i),
			AdminFee:          5000,
			InstallmentAmount: 30000,
			InterestAmount:    2000,
			AssetName:         "Motor",
		}
		err := db.Create(transaction).Error
		assert.Nil(t, err)
	}
}

func CreateAddresses(t *testing.T, contact *entity.Contact, total int) {
	for i := 0; i < total; i++ {
		address := &entity.Address{
			ID:         uuid.NewString(),
			ContactId:  contact.ID,
			Street:     "Jalan Belum Jadi",
			City:       "Jakarta",
			Province:   "DKI Jakarta",
			PostalCode: "2131323",
			Country:    "Indonesia",
		}
		err := db.Create(address).Error
		assert.Nil(t, err)
	}
}

func GetFirstUser(t *testing.T) *entity.User {
	user := new(entity.User)
	err := db.First(user).Error
	assert.Nil(t, err)
	return user
}

func GetFirstContact(t *testing.T, user *entity.User) *entity.Contact {
	contact := new(entity.Contact)
	err := db.Where("user_id = ?", user.ID).First(contact).Error
	assert.Nil(t, err)
	return contact
}

func GetFirstContactLimit(t *testing.T, contact *entity.Contact) *entity.ContactLimit {
	limit := new(entity.ContactLimit)
	err := db.Where("contact_id = ?", contact.ID).First(limit).Error
	assert.Nil(t, err)
	return limit
}

func GetFirstTransaction(t *testing.T, contact *entity.Contact) *entity.Transaction {
	transaction := new(entity.Transaction)
	err := db.Where("contact_id = ?", contact.ID).First(transaction).Error
	assert.Nil(t, err)
	return transaction
}

func GetFirstAddress(t *testing.T, contact *entity.Contact) *entity.Address {
	address := new(entity.Address)
	err := db.Where("contact_id = ?", contact.ID).First(address).Error
	assert.Nil(t, err)
	return address
}
