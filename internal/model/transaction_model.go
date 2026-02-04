package model

type TransactionResponse struct {
	ID                string `json:"id"`
	ContractNumber    string `json:"contract_number"`
	Tenor             int    `json:"tenor"`
	Channel           string `json:"channel"`
	Otr               int64  `json:"otr"`
	AdminFee          int64  `json:"admin_fee"`
	InstallmentAmount int64  `json:"installment_amount"`
	InterestAmount    int64  `json:"interest_amount"`
	AssetName         string `json:"asset_name"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}

type ListTransactionRequest struct {
	UserId    int64  `json:"-" validate:"required"`
	ContactId string `json:"-" validate:"required,max=100,uuid"`
}

type CreateTransactionRequest struct {
	UserId            int64  `json:"-" validate:"required"`
	ContactId         string `json:"-" validate:"required,max=100,uuid"`
	ContractNumber    string `json:"contract_number" validate:"required,max=100"`
	Tenor             int    `json:"tenor" validate:"required,oneof=1 2 3 6"`
	Channel           string `json:"channel" validate:"required,max=50"`
	Otr               int64  `json:"otr" validate:"required,min=0"`
	AdminFee          int64  `json:"admin_fee" validate:"required,min=0"`
	InstallmentAmount int64  `json:"installment_amount" validate:"required,min=0"`
	InterestAmount    int64  `json:"interest_amount" validate:"required,min=0"`
	AssetName         string `json:"asset_name" validate:"required,max=150"`
}

type UpdateTransactionRequest struct {
	UserId            int64  `json:"-" validate:"required"`
	ContactId         string `json:"-" validate:"required,max=100,uuid"`
	ID                string `json:"-" validate:"required,max=100,uuid"`
	ContractNumber    string `json:"contract_number" validate:"required,max=100"`
	Tenor             int    `json:"tenor" validate:"required,oneof=1 2 3 6"`
	Channel           string `json:"channel" validate:"required,max=50"`
	Otr               int64  `json:"otr" validate:"required,min=0"`
	AdminFee          int64  `json:"admin_fee" validate:"required,min=0"`
	InstallmentAmount int64  `json:"installment_amount" validate:"required,min=0"`
	InterestAmount    int64  `json:"interest_amount" validate:"required,min=0"`
	AssetName         string `json:"asset_name" validate:"required,max=150"`
}

type GetTransactionRequest struct {
	UserId    int64  `json:"-" validate:"required"`
	ContactId string `json:"-" validate:"required,max=100,uuid"`
	ID        string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteTransactionRequest struct {
	UserId    int64  `json:"-" validate:"required"`
	ContactId string `json:"-" validate:"required,max=100,uuid"`
	ID        string `json:"-" validate:"required,max=100,uuid"`
}
