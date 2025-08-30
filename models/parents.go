package models

import (
	"time"

	"gorm.io/gorm"
)

type Parent struct {
	ID        uint64      `gorm:"primaryKey"`
	FirstName string    `gorm:"size:100;not null"`
	LastName  string    `gorm:"size:100;not null"`
	Sexe      string    `gorm:"size:10"`  // "M" ou "F"
	Grade     string    `gorm:"size:50"`  // Si tu veux indiquer un statut ou niveau d’instruction
	Position  string    `gorm:"size:100"` // Ex: "Fonctionnaire", "Commerçant"
	Email     string    `gorm:"size:20"`
	Phone     string    `gorm:"size:20"`
	Start     time.Time `gorm:"not null"` 
	Children  int       `gorm:"default:0"` // Nombre d'enfants scolarisés
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateParent(db *gorm.DB, parent *Parent) error {
	return db.Create(parent).Error
}

func GetParentByID(db *gorm.DB, id uint64) (*Parent, error) {
	var parent Parent
	err := db.First(&parent, id).Error
	if err != nil {
		return nil, err
	}
	return &parent, nil
}

func ListParents(db *gorm.DB) ([]Parent, error) {
	var parents []Parent
	err := db.Order("last_name ASC").Find(&parents).Error
	return parents, err
}

func UpdateParent(db *gorm.DB, parent *Parent) error {
	return db.Save(parent).Error
}

func DeleteParent(db *gorm.DB, id uint64) error {
	return db.Delete(&Parent{}, id).Error
}

// Nombre de parent existant dans la database
func CountParent(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Parent{}).Count(&count).Error
	return count, err
}
