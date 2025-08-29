package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID        uint    `gorm:"primaryKey"`
	StudentID uint64    `gorm:"not null"`             // Clé étrangère vers Student
	Student   Student `gorm:"foreignKey:StudentID"` // Relation GORM
	Price     float64 `gorm:"not null"`             // Montant payé
	Type      string  `gorm:"size:50;not null"`     // Ex: "Inscription", "Mensualité"
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CreatePayment enregistre un paiement
func CreatePayment(db *gorm.DB, payment *Payment) error {
	return db.Create(payment).Error
}

// GetPaymentByID récupère un paiement par ID
func GetPaymentByID(db *gorm.DB, id uint) (*Payment, error) {
	var payment Payment
	err := db.Preload("Student").First(&payment, id).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

// ListPayments récupère tous les paiements
func ListPayments(db *gorm.DB) ([]Payment, error) {
	var payments []Payment
	err := db.Preload("Student").Order("created_at DESC").Find(&payments).Error
	return payments, err
}

// UpdatePayment met à jour un paiement
func UpdatePayment(db *gorm.DB, payment *Payment) error {
	return db.Save(payment).Error
}

// DeletePayment supprime un paiement (soft delete)
func DeletePayment(db *gorm.DB, id uint) error {
	return db.Delete(&Payment{}, id).Error
}
