package models

import (
	"time"

	"gorm.io/gorm"
)

type Year struct {
	ID        uint      `gorm:"primaryKey"`
	Label     string    `gorm:"size:20;unique;not null"`
	Start     time.Time `gorm:"not null"`
	End       time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CreateYear crée une nouvelle année scolaire
func CreateYear(db *gorm.DB, year *Year) error {
	return db.Create(year).Error
}

// GetYearByID récupère une année par son ID
func GetYearByID(db *gorm.DB, id uint) (*Year, error) {
	var year Year
	err := db.First(&year, id).Error
	if err != nil {
		return nil, err
	}
	return &year, nil
}

// UpdateYear modifie une année scolaire
func UpdateYear(db *gorm.DB, year *Year) error {
	return db.Save(year).Error
}

// DeleteYear supprime une année scolaire (soft delete si DeletedAt est activé)
func DeleteYear(db *gorm.DB, id uint) error {
	return db.Delete(&Year{}, id).Error
}

// ListYears liste toutes les années scolaires
func ListYears(db *gorm.DB) ([]Year, error) {
	var years []Year
	err := db.Order("start DESC").Find(&years).Error
	return years, err
}
