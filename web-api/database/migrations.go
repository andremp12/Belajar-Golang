package database

import "time"

type Book struct {
	ID          int    `gorm:"type:int(11);primary_key;auto_increment"`
	Title       string `gorm:"type:varchar(50);not null"`
	Price       int    `gorm:"type:bigint;not null"`
	Rating      int    `gorm:"type:int(2);null"`
	Description string `gorm:"type:text;null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
