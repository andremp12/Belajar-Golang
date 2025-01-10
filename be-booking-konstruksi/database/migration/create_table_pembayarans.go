package database

import "time"

type Pembayaran struct {
	ID           int        `gorm:"type:int(11);primary_key;auto_increment"`
	KonstruksiID int        `gorm:"type:int(11);not null"`
	Konstruksis  Konstruksi `gorm:"references:ID;foreignKey:KonstruksiID"`
	Name         string     `gorm:"type:varchar(10);not null"`
	Kode         string     `gorm:"type:varchar(50);not null;unique"`
	Metode       string     `gorm:"type:varchar(50);not null"`
	Status       string     `gorm:"type:varchar(20);not null"`
	Total        int        `gorm:"type:bigInt;not null"`
	CreatedAt    time.Time  `gorm:"type:timestamp;not null"`
	UpdatedAt    time.Time  `gorm:"type:timestamp;not null"`
}
