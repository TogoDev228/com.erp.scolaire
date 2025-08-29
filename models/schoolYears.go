package models

import (
	"time"

	"gorm.io/gorm"
)

type SchoolYear struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"size:20;unique;not null"`
	StartYear time.Time `gorm:"not null"`
	EndYear   time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Crée une nouvelle année scolaire
func CreateSchoolYear(db *gorm.DB, schoolYear *SchoolYear) error {
	return db.Create(schoolYear).Error
}

// Récupère une année par son ID
func GetSchoolYearByID(db *gorm.DB, id uint) (*SchoolYear, error) {
	var schoolYear SchoolYear
	err := db.First(&schoolYear, id).Error
	if err != nil {
		return nil, err
	}
	return &schoolYear, nil
}

// Modifie une année scolaire
func UpdateSchoolYear(db *gorm.DB, schoolYear *SchoolYear) error {
	return db.Save(schoolYear).Error
}

// Supprime une année scolaire (soft delete si DeletedAt est activé)
func DeleteSchoolYear(db *gorm.DB, id uint) error {
	return db.Delete(&SchoolYear{}, id).Error
}

// Liste toutes les années scolaires
func ListSchoolYears(db *gorm.DB) ([]SchoolYear, error) {
	var schoolYear []SchoolYear
	err := db.Order("start DESC").Find(&schoolYear).Error
	return schoolYear, err
}
