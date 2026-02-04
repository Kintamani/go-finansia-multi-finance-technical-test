package converter

import (
	"go-finansia-multi-finance-technical-test/internal/entity"
	"go-finansia-multi-finance-technical-test/internal/model"
	"time"
)

func ContactLimitToResponse(limit *entity.ContactLimit, usedAmount int64) *model.ContactLimitResponse {
	remaining := limit.LimitAmount - usedAmount
	if remaining < 0 {
		remaining = 0
	}

	return &model.ContactLimitResponse{
		ID:              limit.ID,
		Tenor:           limit.Tenor,
		LimitAmount:     limit.LimitAmount,
		UsedAmount:      usedAmount,
		RemainingAmount: remaining,
		CreatedAt:       limit.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       limit.UpdatedAt.Format(time.RFC3339),
	}
}
