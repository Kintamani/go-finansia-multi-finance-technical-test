package model

type ContactLimitResponse struct {
	ID              string `json:"id"`
	Tenor           int    `json:"tenor"`
	LimitAmount     int64  `json:"limit_amount"`
	UsedAmount      int64  `json:"used_amount"`
	RemainingAmount int64  `json:"remaining_amount"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

type ListContactLimitRequest struct {
	UserId    int64  `json:"-" validate:"required"`
	ContactId string `json:"-" validate:"required,max=100,uuid"`
}

type CreateContactLimitRequest struct {
	UserId      int64  `json:"-" validate:"required"`
	ContactId   string `json:"-" validate:"required,max=100,uuid"`
	Tenor       int    `json:"tenor" validate:"required,oneof=1 2 3 6"`
	LimitAmount int64  `json:"limit_amount" validate:"required,min=0"`
}

type UpdateContactLimitRequest struct {
	UserId      int64  `json:"-" validate:"required"`
	ContactId   string `json:"-" validate:"required,max=100,uuid"`
	ID          string `json:"-" validate:"required,max=100,uuid"`
	Tenor       int    `json:"tenor" validate:"required,oneof=1 2 3 6"`
	LimitAmount int64  `json:"limit_amount" validate:"required,min=0"`
}

type GetContactLimitRequest struct {
	UserId    int64  `json:"-" validate:"required"`
	ContactId string `json:"-" validate:"required,max=100,uuid"`
	ID        string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteContactLimitRequest struct {
	UserId    int64  `json:"-" validate:"required"`
	ContactId string `json:"-" validate:"required,max=100,uuid"`
	ID        string `json:"-" validate:"required,max=100,uuid"`
}
