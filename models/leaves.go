package models

import (
	"time"

	"gorm.io/gorm"
)

type Leave struct {
	ID        uint64    `gorm:"primaryKey"`
	UserID    uint64    `gorm:"not null"`
	UserType  string    `gorm:"size:50;not null"`
	Status    string    `gorm:"size:50;not null"`
	StartYear time.Time `gorm:"not null"`
	EndYear   time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Ajoute un nouvau conger pour un utilisateur donnée
func CreateLeave(db *gorm.DB, leave *Leave) error {
	return db.Create(leave).Error
}

// Récupère un conger via son ID
func GetLeaveByID(db *gorm.DB, id uint) (*Leave, error) {
	var leave Leave
	err := db.First(&leave, id).Error
	if err != nil {
		return nil, err
	}
	return &leave, nil
}

// Modifie un conger existant
func UpdateLeave(db *gorm.DB, leave *Leave) error {
	return db.Save(leave).Error
}

// Supprime un conger (soft delete si activé)
func DeleteLeave(db *gorm.DB, id uint) error {
	return db.Delete(&Leave{}, id).Error
}

// Retourne tous les conger (optionnel : avec pagination ou filtre par année)
func ListLeaves(db *gorm.DB) ([]Leave, error) {
	var leaves []Leave
	err := db.Order("last_name ASC").Find(&leaves).Error
	return leaves, err
}

// Nombre de conger existant dans la database
func CountLeave(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Leave{}).Count(&count).Error
	return count, err
}
