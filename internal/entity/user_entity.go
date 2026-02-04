package entity

import "time"

// User is a struct that represents a user entity
type User struct {
	ID                  int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Username            string    `gorm:"column:username"`
	Password            string    `gorm:"column:password"`
	Name                string    `gorm:"column:name"`
	Token               string    `gorm:"column:token"`
	FailedLoginAttempts int       `gorm:"column:failed_login_attempts"`
	LastFailedLoginAt   int64     `gorm:"column:last_failed_login_at"`
	LockedUntil         int64     `gorm:"column:locked_until"`
	LastLoginAt         int64     `gorm:"column:last_login_at"`
	CreatedAt           time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt           time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Contacts            []Contact `gorm:"foreignKey:user_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}
