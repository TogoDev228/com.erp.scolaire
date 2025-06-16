package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:100;not null"` // Nom de l’activité
	Description string    `gorm:"type:text"`         // Détails
	Type        string    `gorm:"size:50"`           // Ex : "Sportive", "Culturelle", etc.
	Budget      string    `gorm:"size:100"`          // Montant alloué
	Location    string    `gorm:"size:100"`          // Lieu
	Start       time.Time `gorm:"not null"`          // Date de début
	End         time.Time `gorm:"not null"`          // Date de fin
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

// Nombre de prof existant dans la database
func CountActivities(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Activity{}).Count(&count).Error
	return count, err
}
