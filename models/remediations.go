package models

import (
	"time"

	"gorm.io/gorm"
)

type Remediation struct {
	ID          uint64   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text"`
	StudentID   uint64 `gorm:"not null"`
	TeacherID   uint64 `gorm:"not null"`
	Status      string `gorm:"size:50;not null"`
	Start       time.Time
	End         time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// CreatePayment enregistre un paiement
func CreateRemediation(db *gorm.DB, remediation *Remediation) error {
	return db.Create(remediation).Error
}

// GetPaymentByID récupère un paiement par ID
func GetRemediationByID(db *gorm.DB, id uint) (*Remediation, error) {
	var remediation Remediation
	err := db.First(&remediation, id).Error
	if err != nil {
		return nil, err
	}
	return &remediation, nil
}

// ListPayments récupère tous les paiements
func ListRemediations(db *gorm.DB) ([]Remediation, error) {
	var remediations []Remediation
	err := db.Order("created_at DESC").Find(&remediations).Error
	return remediations, err
}

// UpdatePayment met à jour un paiement
func UpdateRemediation(db *gorm.DB, remediation *Remediation) error {
	return db.Save(remediation).Error
}

// DeletePayment supprime un paiement (soft delete)
func DeleteRemediation(db *gorm.DB, id uint) error {
	return db.Delete(&Remediation{}, id).Error
}
