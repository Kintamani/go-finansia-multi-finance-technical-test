package converter

import (
	"go-finansia-multi-finance-technical-test/internal/entity"
	"go-finansia-multi-finance-technical-test/internal/model"
	"time"
)

func TransactionToResponse(transaction *entity.Transaction) *model.TransactionResponse {
	return &model.TransactionResponse{
		ID:                transaction.ID,
		ContractNumber:    transaction.ContractNumber,
		Tenor:             transaction.Tenor,
		Channel:           transaction.Channel,
		Otr:               transaction.Otr,
		AdminFee:          transaction.AdminFee,
		InstallmentAmount: transaction.InstallmentAmount,
		InterestAmount:    transaction.InterestAmount,
		AssetName:         transaction.AssetName,
		CreatedAt:         transaction.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         transaction.UpdatedAt.Format(time.RFC3339),
	}
}
