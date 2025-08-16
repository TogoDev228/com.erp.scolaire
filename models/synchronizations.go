package models

import (
	"time"

	"gorm.io/gorm"
)

type Synchronization struct {
	ID        uint   `gorm:"primaryKey"`
	Type     string `gorm:"not null"`
	Status     string `gorm:"size:50;not null"`
	Start     time.Time `gorm:"not null"`
	End     time.Time `gorm:"not null"`
	AttemptCount     int `gorm:"default:0"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CreateClass ajoute une nouvelle classe
func CreateSynchronization(db *gorm.DB, synchronization *Synchronization) error {
	return db.Create(synchronization).Error
}

// GetClassByID récupère une classe par ID
func GetSynchronizationByID(db *gorm.DB, id uint) (*Synchronization, error) {
	var synchronization Synchronization
	err := db.First(&synchronization, id).Error
	if err != nil {
		return nil, err
	}
	return &synchronization, nil
}

// ListClasses retourne toutes les classes (ordre alphabétique)
func ListSynchronizations(db *gorm.DB) ([]Synchronization, error) {
	var synchronizations []Synchronization
	err := db.Order("title ASC").Find(&synchronizations).Error
	return synchronizations, err
}

// DeleteClass supprime une classe (soft delete)
func DeleteSynchronization(db *gorm.DB, id uint) error {
	return db.Delete(&Synchronization{}, id).Error
}

// Nombre de Class existant dans la database
func CountSynchronization(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Synchronization{}).Count(&count).Error
	return count, err
}
