package model

type ContactResponse struct {
	ID          string            `json:"id"`
	NIK         string            `json:"nik"`
	FullName    string            `json:"full_name"`
	LegalName   string            `json:"legal_name"`
	BirthPlace  string            `json:"birth_place"`
	BirthDate   string            `json:"birth_date"`
	Salary      int64             `json:"salary"`
	KtpPhoto    string            `json:"ktp_photo"`
	SelfiePhoto string            `json:"selfie_photo"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	Addresses   []AddressResponse `json:"addresses,omitempty"`
}

type CreateContactRequest struct {
	UserId      int64  `json:"-" validate:"required"`
	NIK         string `json:"nik" validate:"required,len=16,numeric"`
	FullName    string `json:"full_name" validate:"required,max=150"`
	LegalName   string `json:"legal_name" validate:"required,max=150"`
	BirthPlace  string `json:"birth_place" validate:"required,max=100"`
	BirthDate   string `json:"birth_date" validate:"required"`
	Salary      int64  `json:"salary" validate:"required,min=0"`
	KtpPhoto    string `json:"ktp_photo" validate:"required,max=500"`
	SelfiePhoto string `json:"selfie_photo" validate:"required,max=500"`
}

type UpdateContactRequest struct {
	UserId      int64  `json:"-" validate:"required"`
	ID          string `json:"-" validate:"required,max=100,uuid"`
	NIK         string `json:"nik" validate:"required,len=16,numeric"`
	FullName    string `json:"full_name" validate:"required,max=150"`
	LegalName   string `json:"legal_name" validate:"required,max=150"`
	BirthPlace  string `json:"birth_place" validate:"required,max=100"`
	BirthDate   string `json:"birth_date" validate:"required"`
	Salary      int64  `json:"salary" validate:"required,min=0"`
	KtpPhoto    string `json:"ktp_photo" validate:"required,max=500"`
	SelfiePhoto string `json:"selfie_photo" validate:"required,max=500"`
}

type SearchContactRequest struct {
	UserId    int64  `json:"-" validate:"required"`
	NIK       string `json:"nik" validate:"max=16"`
	FullName  string `json:"full_name" validate:"max=150"`
	LegalName string `json:"legal_name" validate:"max=150"`
	Page      int    `json:"page" validate:"min=1"`
	Size      int    `json:"size" validate:"min=1,max=100"`
}

type GetContactRequest struct {
	UserId int64  `json:"-" validate:"required"`
	ID     string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteContactRequest struct {
	UserId int64  `json:"-" validate:"required"`
	ID     string `json:"-" validate:"required,max=100,uuid"`
}
