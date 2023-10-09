package entity

import "time"

type User struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time

	NIM      string `gorm:"type:VARCHAR(20);UNIQUE NOT NULL"`
	Password string `gorm:"type:VARCHAR(255);NOT NULL"`
	Nama         string `gorm:"type:VARCHAR(255)"`
	Jenjang      string `gorm:"type:VARCHAR(255)"`
	Fakultas     string `gorm:"type:VARCHAR(255)"`
	Jurusan      string `gorm:"type:VARCHAR(255)"`
	ProgramStudi string `gorm:"type:VARCHAR(255)"`
	Seleksi      string `gorm:"type:VARCHAR(255)"`
	NomorUjian   string `gorm:"type:VARCHAR(255)"`
	FotoProfil   string `gorm:"type:VARCHAR(255)"`
}