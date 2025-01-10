package database

import "time"

type Konstruksi struct {
	ID               int            `gorm:"type:int(11);primary_key;auto_increment"`
	TipeKonstruksiID int            `gorm:"type:int(11);not null"`
	TipeKonstruksi   TipeKonstruksi `gorm:"references:ID"`
	UserID           int            `gorm:"type:int(11);not null"`
	User             User           `gorm:"references:ID"`
	Tipe             string         `gorm:"type:varchar(50);not null"`
	Address          string         `gorm:"type:text;not null"`
	Total            int            `gorm:"type:bigInt;not null"`
	StartDate        time.Time      `gorm:"type:date; null"`
	EndDate          time.Time      `gorm:"type:date; null"`
	CreatedAt        time.Time      `gorm:"type:timestamp;not null"`
	UpdatedAt        time.Time      `gorm:"type:timestamp;not null"`
}
