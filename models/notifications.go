package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:100;not null;unique"`
	Description string `gorm:"size:100;not null;unique"`
	Status      string `gorm:"size:100;not null;unique"`
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// Ajoute une nouvelle notification
func CreateNotification(db *gorm.DB, notification *Notification) error {
	return db.Create(notification).Error
}

// Récupère une notification par ID
func GetNotificationByID(db *gorm.DB, id uint) (*Notification, error) {
	var notification Notification
	err := db.First(&notification, id).Error
	if err != nil {
		return nil, err
	}
	return &notification, nil
}

// Retourne toutes les notification (ordre alphabétique)
func ListNotification(db *gorm.DB) ([]Notification, error) {
	var notifications []Notification
	err := db.Order("title ASC").Find(&notifications).Error
	return notifications, err
}

// Supprime une notification (soft delete)
func DeleteNotification(db *gorm.DB, id uint) error {
	return db.Delete(&Notification{}, id).Error
}

// Nombre de notification existant dans la database
func CountNotification(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Notification{}).Count(&count).Error
	return count, err
}
