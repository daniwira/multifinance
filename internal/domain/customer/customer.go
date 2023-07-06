package customer

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID           uint64         `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	Nik          string         `gorm:"column:nik;unique;size:16" json:"nik"`
	FullName     string         `gorm:"column:full_name" json:"full_name"`
	LegalName    string         `gorm:"column:legal_name" json:"legal_name"`
	PlaceOfBirth string         `gorm:"column:place_of_birth" json:"place_of_birth"`
	DateOfBirth  string         `gorm:"column:date_of_birth" json:"date_of_birth"`
	Salary       float64        `gorm:"column:salary" json:"salary"`
	KTPPhoto     string         `gorm:"column:ktp_photo" json:"ktp_photo"`
	SelfiePhoto  string         `gorm:"column:selfie_photo" json:"selfie_photo"`
}
