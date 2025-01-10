package database

import "time"

type Laporan struct {
	ID           int        `gorm:"type:int(11);primary_key;auto_increment"`
	KonstruksiID int        `gorm:"type:int(11);not null"`
	Konstruksis  Konstruksi `gorm:"references:ID;foreignKey:KonstruksiID"`
	UserID       int        `gorm:"type:int(11);not null"`
	User         User       `gorm:"references:ID;foreignKey:UserID"`
	Name         string     `gorm:"type:varchar(100);not null"`
	Tipe         string     `gorm:"type:varchar(50);not null"`
	Pelanggan    string     `gorm:"type:varchar(50);not null"`
	Mandor       string     `gorm:"type:varchar(50);not null"`
	Deskripsi    string     `gorm:"type:text; null"`
	CreatedAt    time.Time  `gorm:"type:timestamp;not null"`
	UpdatedAt    time.Time  `gorm:"type:timestamp;not null"`
}
