package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:100;not null"`
	Description string    `gorm:"type:text"`
	Type        string    `gorm:"size:50"`
	Budget      string    `gorm:"size:100"`
	Location    string    `gorm:"size:100"`
	Status      string    `gorm:"size:100"`
	SchoolYear  string    `gorm:"size:100"`
	Creator     string    `gorm:"size:120;not null"`
	Start       time.Time `gorm:"not null"`
	End         time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// CreateActivity ajoute une nouvelle activité
func CreateActivity(db *gorm.DB, activity *Activity) error {
	return db.Create(activity).Error
}

// GetActivityByID récupère une activité par ID
func GetActivityByID(db *gorm.DB, id uint) (*Activity, error) {
	var activity Activity
	err := db.First(&activity, id).Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

// ListActivities retourne toutes les activités
func ListActivities(db *gorm.DB) ([]Activity, error) {
	var activities []Activity
	err := db.Order("start DESC").Find(&activities).Error
	return activities, err
}

// UpdateActivity met à jour une activité
func UpdateActivity(db *gorm.DB, activity *Activity) error {
	return db.Save(activity).Error
}

// DeleteActivity supprime une activité (soft delete)
func DeleteActivity(db *gorm.DB, id uint) error {
	return db.Delete(&Activity{}, id).Error
}

// Nombre de Activity existant dans la database
func CountActivities(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Activity{}).Count(&count).Error
	return count, err
}
