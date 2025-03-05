package models

import "time"

type AuthUser struct{
	ID uint `gorm:"primaryKey;autoIncrement"`
	UserID int
	User User
	Token string `gorm:"unique;not null"`
	CreatedAt time.Time
	ExpiresAt time.Time `gorm:"not null"`
}