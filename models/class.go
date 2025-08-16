package models

import (
	"time"

	"gorm.io/gorm"
)

type Class struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:100;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CreateClass ajoute une nouvelle classe
func CreateClass(db *gorm.DB, class *Class) error {
	return db.Create(class).Error
}

// GetClassByID récupère une classe par ID
func GetClassByID(db *gorm.DB, id uint) (*Class, error) {
	var class Class
	err := db.First(&class, id).Error
	if err != nil {
		return nil, err
	}
	return &class, nil
}

// ListClasses retourne toutes les classes (ordre alphabétique)
func ListClasses(db *gorm.DB) ([]Class, error) {
	var classes []Class
	err := db.Order("title ASC").Find(&classes).Error
	return classes, err
}

// UpdateClass met à jour une classe
func UpdateClass(db *gorm.DB, class *Class) error {
	return db.Save(class).Error
}

// DeleteClass supprime une classe (soft delete)
func DeleteClass(db *gorm.DB, id uint) error {
	return db.Delete(&Class{}, id).Error
}

// Nombre de Class existant dans la database
func CountClass(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Class{}).Count(&count).Error
	return count, err
}
