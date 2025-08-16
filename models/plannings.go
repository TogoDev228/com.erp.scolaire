package models

import (
	"time"

	"gorm.io/gorm"
)

type Planning struct {
	ID         uint      `gorm:"primaryKey"`
	Title      int       `gorm:"size:120;unique;not null"`
	Type       string    `gorm:"size:70;not null"`
	Class      string    `gorm:"size:100;not null"`
	SchoolYear string    `gorm:"not null"`
	Start      time.Time `gorm:"not null"`
	End        time.Time `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// Ajoute un nouvau conger pour un utilisateur donnée
func CreatePlanning(db *gorm.DB, planning *Planning) error {
	return db.Create(planning).Error
}

// Récupère un conger via son ID
func GetPlanningByID(db *gorm.DB, id uint) (*Planning, error) {
	var planning Planning
	err := db.First(&planning, id).Error
	if err != nil {
		return nil, err
	}
	return &planning, nil
}

// Modifie un conger existant
func UpdatePlanning(db *gorm.DB, planning *Planning) error {
	return db.Save(planning).Error
}

// Supprime un conger (soft delete si activé)
func DeletePlanning(db *gorm.DB, id uint) error {
	return db.Delete(&Planning{}, id).Error
}

// Retourne tous les conger (optionnel : avec pagination ou filtre par année)
func ListPlannings(db *gorm.DB) ([]Planning, error) {
	var plannings []Planning
	err := db.Order("last_name ASC").Find(&plannings).Error
	return plannings, err
}

// Nombre de conger existant dans la database
func CountPlanning(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Planning{}).Count(&count).Error
	return count, err
}
