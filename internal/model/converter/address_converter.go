package converter

import (
	"go-finansia-multi-finance-technical-test/internal/entity"
	"go-finansia-multi-finance-technical-test/internal/model"
	"time"
)

func AddressToResponse(address *entity.Address) *model.AddressResponse {
	return &model.AddressResponse{
		ID:         address.ID,
		Street:     address.Street,
		City:       address.City,
		Province:   address.Province,
		PostalCode: address.PostalCode,
		Country:    address.Country,
		CreatedAt:  address.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  address.UpdatedAt.Format(time.RFC3339),
	}
}
