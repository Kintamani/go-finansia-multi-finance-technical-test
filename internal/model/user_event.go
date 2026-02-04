package model

import "strconv"

type UserEvent struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func (u *UserEvent) GetId() string {
	return strconv.FormatInt(u.ID, 10)
}
