package models

import (
	"time"

	"gorm.io/gorm"
)

type Remuneration struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	UserType  string  `gorm:"size:50;not null"`
	Price     float64 `gorm:"not null"`
	Type      string  `gorm:"size:50;not null"`
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Enregistre un paiement de rémunération
func CreateRemuneration(db *gorm.DB, remuneration *Remuneration) error {
	return db.Create(remuneration).Error
}

// Récupère un paiement par ID
func GetRemunerationByID(db *gorm.DB, id uint) (*Remuneration, error) {
	var remuneration Remuneration
	err := db.First(&remuneration, id).Error
	if err != nil {
		return nil, err
	}
	return &remuneration, nil
}

// Récupère tous les paiements
func ListRemuneration(db *gorm.DB) ([]Remuneration, error) {
	var remunerations []Remuneration
	err := db.Order("created_at DESC").Find(&remunerations).Error
	return remunerations, err
}

// Met à jour un paiement
func UpdateRemuneration(db *gorm.DB, remuneration *Remuneration) error {
	return db.Save(remuneration).Error
}

// Supprime un paiement (soft delete)
func DeleteRemuneration(db *gorm.DB, id uint) error {
	return db.Delete(&Remuneration{}, id).Error
}
