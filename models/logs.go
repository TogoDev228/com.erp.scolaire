package models

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	ID        uint   `gorm:"primaryKey"`
	Type      string `gorm:"size:50;not null"`   // Ex: "INFO", "ERROR", "LOGIN", etc.
	Message   string `gorm:"type:text;not null"` // Contenu du log
	UserID    uint   `gorm:"not null"`           // FK vers l'utilisateur concerné
	User      User   `gorm:"foreignKey:UserID"`  // Relation avec User
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CreateLog enregistre une nouvelle entrée de log
func CreateLog(db *gorm.DB, log *Log) error {
	return db.Create(log).Error
}

// GetLogByID récupère un log par son ID
func GetLogByID(db *gorm.DB, id uint) (*Log, error) {
	var log Log
	err := db.Preload("User").First(&log, id).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// ListLogs retourne tous les logs (avec tri du plus récent)
func ListLogs(db *gorm.DB) ([]Log, error) {
	var logs []Log
	err := db.Preload("User").Order("created_at DESC").Find(&logs).Error
	return logs, err
}

// DeleteLog supprime un log (soft delete)
func DeleteLog(db *gorm.DB, id uint) error {
	return db.Delete(&Log{}, id).Error
}

// Nombre de Log existant dans la database
func CountLogs(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Log{}).Count(&count).Error
	return count, err
}
