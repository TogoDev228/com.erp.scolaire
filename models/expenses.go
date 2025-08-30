package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	ID          uint64   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text"`
	Type        string `gorm:"not null"`
	Status      string `gorm:"not null"`
	SchoolYear  string `gorm:"size:50;not null"`
	Creator     string `gorm:"size:150;not null"`
	Start       time.Time
	End         time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// CreatePayment enregistre un paiement
func CreateExpense(db *gorm.DB, expense *Expense) error {
	return db.Create(expense).Error
}

// GetPaymentByID récupère un paiement par ID
func GetExpenseByID(db *gorm.DB, id uint64) (*Expense, error) {
	var expense Expense
	err := db.First(&expense, id).Error
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

// ListPayments récupère tous les paiements
func ListExpenses(db *gorm.DB) ([]Expense, error) {
	var expenses []Expense
	err := db.Order("created_at DESC").Find(&expenses).Error
	return expenses, err
}

// UpdatePayment met à jour un paiement
func UpdateExpense(db *gorm.DB, expense *Expense) error {
	return db.Save(expense).Error
}

// DeletePayment supprime un paiement (soft delete)
func DeleteExpense(db *gorm.DB, id uint64) error {
	return db.Delete(&Expense{}, id).Error
}
