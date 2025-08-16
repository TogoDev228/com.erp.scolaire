package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint   `gorm:"primaryKey"`
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
func CreateTask(db *gorm.DB, task *Task) error {
	return db.Create(task).Error
}

// GetPaymentByID récupère un paiement par ID
func GetTaskByID(db *gorm.DB, id uint) (*Task, error) {
	var task Task
	err := db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// ListPayments récupère tous les paiements
func ListTasks(db *gorm.DB) ([]Task, error) {
	var tasks []Task
	err := db.Order("created_at DESC").Find(&tasks).Error
	return tasks, err
}

// UpdatePayment met à jour un paiement
func UpdateTask(db *gorm.DB, task *Task) error {
	return db.Save(task).Error
}

// DeletePayment supprime un paiement (soft delete)
func DeleteTask(db *gorm.DB, id uint) error {
	return db.Delete(&Task{}, id).Error
}
