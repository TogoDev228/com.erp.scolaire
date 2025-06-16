package models

import (
	"time"

	"gorm.io/gorm"
)

type History struct {
	ID        uint    `gorm:"primaryKey"`
	Title     string  `gorm:"size:255;not null"` // Intitulé de l’opération ou événement
	Price     float64 `gorm:"not null"`          // Montant associé (facultatif selon le contexte)
	Type      string  `gorm:"size:50"`           // Ex: "Paiement", "Achat", "Réparation"
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete activé
}

// CreateHistory ajoute une entrée dans l’historique
func CreateHistory(db *gorm.DB, history *History) error {
	return db.Create(history).Error
}

// GetHistoryByID récupère une entrée de l’historique par ID
func GetHistoryByID(db *gorm.DB, id uint) (*History, error) {
	var history History
	err := db.First(&history, id).Error
	if err != nil {
		return nil, err
	}
	return &history, nil
}

// ListHistories retourne toutes les entrées de l’historique
func ListHistories(db *gorm.DB) ([]History, error) {
	var histories []History
	err := db.Order("created_at DESC").Find(&histories).Error
	return histories, err
}

// UpdateHistory met à jour une entrée
func UpdateHistory(db *gorm.DB, history *History) error {
	return db.Save(history).Error
}

// DeleteHistory supprime une entrée (soft delete)
func DeleteHistory(db *gorm.DB, id uint) error {
	return db.Delete(&History{}, id).Error
}

// Nombre de prof existant dans la database
func CountHistory(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Prof{}).Count(&count).Error
	return count, err
}
