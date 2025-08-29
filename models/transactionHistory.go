package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionHistory struct {
	ID        uint    `gorm:"primaryKey"`
	Title     string  `gorm:"size:200;not null"` // Intitulé de l’opération ou événement
	Price     float64 `gorm:"not null"`          // Montant associé (facultatif selon le contexte)
	Type      string  `gorm:"size:50"`           // Ex: "Paiement", "Achat", "Réparation"
	Actor     string  `gorm:"size:80"`           // Acteur de l’opération (ex: "Client", "Fournisseur")
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete activé
}

// CreateHistory ajoute une entrée dans l’historique
func CreateTransactionHistory(db *gorm.DB, transactionHistory *TransactionHistory) error {
	return db.Create(transactionHistory).Error
}

// GetHistoryByID récupère une entrée de l’historique par ID
func GetTransactionHistoryByID(db *gorm.DB, id uint) (*TransactionHistory, error) {
	var transactionHistory TransactionHistory
	err := db.First(&transactionHistory, id).Error
	if err != nil {
		return nil, err
	}
	return &transactionHistory, nil
}

// ListHistories retourne toutes les entrées de l’historique
func ListTransactionHistories(db *gorm.DB) ([]TransactionHistory, error) {
	var transactionHistories []TransactionHistory
	err := db.Order("created_at DESC").Find(&transactionHistories).Error
	return transactionHistories, err
}

// UpdateHistory met à jour une entrée
func UpdateTransactionHistory(db *gorm.DB, transactionHistory *TransactionHistory) error {
	return db.Save(transactionHistory).Error
}

// DeleteHistory supprime une entrée (soft delete)
func DeleteTransactionHistory(db *gorm.DB, id uint) error {
	return db.Delete(&TransactionHistory{}, id).Error
}

// Nombre de transaction existant dans la database
func CountTransactionHistory(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&TransactionHistory{}).Count(&count).Error
	return count, err
}
