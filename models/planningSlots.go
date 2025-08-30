package models

import (
	"time"

	"gorm.io/gorm"
)

type PlanningSlot struct {
	ID        uint64    `gorm:"primaryKey"`
	PlaningID uint64    `gorm:"not null"`
	Lesson    string    `gorm:"size:150;not null"`
	Teacher   string    `gorm:"size:150;not null"`
	StartHour time.Time `gorm:"type:time;not null"`
	EndHour   time.Time `gorm:"type:time;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Ajoute un nouvau conger pour un utilisateur donnée
func CreatePlanningSlot(db *gorm.DB, planningSlot *PlanningSlot) error {
	return db.Create(planningSlot).Error
}

// Récupère un conger via son ID
func GetPlanningSlotByID(db *gorm.DB, id uint64) (*PlanningSlot, error) {
	var planningSlot PlanningSlot
	err := db.First(&planningSlot, id).Error
	if err != nil {
		return nil, err
	}
	return &planningSlot, nil
}

// Modifie un conger existant
func UpdatePlanningSlot(db *gorm.DB, planningSlot *PlanningSlot) error {
	return db.Save(planningSlot).Error
}

// Supprime un conger (soft delete si activé)
func DeletePlanningSlot(db *gorm.DB, id uint64) error {
	return db.Delete(&PlanningSlot{}, id).Error
}

// Retourne tous les conger (optionnel : avec pagination ou filtre par année)
func ListPlanningSlots(db *gorm.DB) ([]PlanningSlot, error) {
	var planningSlots []PlanningSlot
	err := db.Find(&planningSlots).Error
	return planningSlots, err
}

// Nombre de conger existant dans la database
func CountPlanningSlot(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&PlanningSlot{}).Count(&count).Error
	return count, err
}
