package entity

import "time"

type User struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time

	NIM      string `gorm:"type:VARCHAR(20);UNIQUE"`
	Password string `gorm:"type:VARCHAR(255);NOT NULL"`
}