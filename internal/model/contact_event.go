package model

type ContactEvent struct {
	ID          string `json:"id"`
	UserID      int64  `json:"user_id"`
	NIK         string `json:"nik"`
	FullName    string `json:"full_name"`
	LegalName   string `json:"legal_name"`
	BirthPlace  string `json:"birth_place"`
	BirthDate   string `json:"birth_date"`
	Salary      int64  `json:"salary"`
	KtpPhoto    string `json:"ktp_photo"`
	SelfiePhoto string `json:"selfie_photo"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (c *ContactEvent) GetId() string {
	return c.ID
}
