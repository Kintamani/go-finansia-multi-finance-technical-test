package converter

import (
	"go-finansia-multi-finance-technical-test/internal/entity"
	"go-finansia-multi-finance-technical-test/internal/model"
	"time"
)

func ContactToResponse(contact *entity.Contact) *model.ContactResponse {
	return &model.ContactResponse{
		ID:          contact.ID,
		NIK:         contact.NIK,
		FullName:    contact.FullName,
		LegalName:   contact.LegalName,
		BirthPlace:  contact.BirthPlace,
		BirthDate:   contact.BirthDate,
		Salary:      contact.Salary,
		KtpPhoto:    contact.KtpPhoto,
		SelfiePhoto: contact.SelfiePhoto,
		CreatedAt:   contact.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   contact.UpdatedAt.Format(time.RFC3339),
	}
}

func ContactToEvent(contact *entity.Contact) *model.ContactEvent {
	return &model.ContactEvent{
		ID:          contact.ID,
		UserID:      contact.UserId,
		NIK:         contact.NIK,
		FullName:    contact.FullName,
		LegalName:   contact.LegalName,
		BirthPlace:  contact.BirthPlace,
		BirthDate:   contact.BirthDate,
		Salary:      contact.Salary,
		KtpPhoto:    contact.KtpPhoto,
		SelfiePhoto: contact.SelfiePhoto,
		CreatedAt:   contact.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   contact.UpdatedAt.Format(time.RFC3339),
	}
}
