package models

import (
	"time"

	"gorm.io/gorm"
)

type StudentClass struct {
	ID         uint64   `gorm:"primaryKey"`
	StudentID  uint64    `gorm:"not null"`
	Class      string `gorm:"not null"`
	Reason      string `gorm:"not null"`
	SchoolYear string `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// Crée une nouvelle relation entre un élève et une classe
func CreateStudentClass(db *gorm.DB, studentClass *StudentClass) error {
	return db.Create(studentClass).Error
}

// Récupère une relation entre un élève et une classe
func GetStudentClassByID(db *gorm.DB, id uint64) (*StudentClass, error) {
	var studentClass StudentClass
	err := db.First(&studentClass, id).Error
	if err != nil {
		return nil, err
	}
	return &studentClass, nil
}

// Modifie une relation entre un élève et une classe
func UpdateStudentClass(db *gorm.DB, studentClass *StudentClass) error {
	return db.Save(studentClass).Error
}

// Supprime une relation entre un élève et une classe
func DeleteStudentClass(db *gorm.DB, id uint64) error {
	return db.Delete(&StudentClass{}, id).Error
}

// Liste toutes les relation entre un élève et des classes
func ListStudentClass(db *gorm.DB) ([]StudentClass, error) {
	var studentClasses []StudentClass
	err := db.Order("start DESC").Find(&studentClasses).Error
	return studentClasses, err
}
