package models

import (
	"time"

	"gorm.io/gorm"
)

type Staff struct {
	ID        uint64    `gorm:"primaryKey"`
	Matricul  string    `gorm:"unique;not null"`
	FirstName string    `gorm:"size:100;not null"`
	LastName  string    `gorm:"size:100;not null"`
	Position  string    `gorm:"size:100;not null"` // Ex: "Secrétaire", "Comptable", etc.
	Sexe      string    `gorm:"size:25"`           // "M", "F", ou autres
	Email     string    `gorm:"size:120;unique"`
	Grade     string    `gorm:"size:50"` // Si pertinent
	Phone     string    `gorm:"size:20"`
	Start     time.Time `gorm:"not null"`
	End       time.Time `gorm:"null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Pour soft delete
}

// CreateStaff ajoute un nouveau membre du personnel
func CreateStaff(db *gorm.DB, staff *Staff) error {
	return db.Create(staff).Error
}

// GetStaffByID récupère un staff par ID
func GetStaffByID(db *gorm.DB, id uint64) (*Staff, error) {
	var staff Staff
	err := db.First(&staff, id).Error
	if err != nil {
		return nil, err
	}
	return &staff, nil
}

// GetStaffByMatricul récupère un staff par matricule
func GetStaffByMatricul(db *gorm.DB, matricul string) (*Staff, error) {
	var staff Staff
	err := db.Where("matricul = ?", matricul).First(&staff).Error
	if err != nil {
		return nil, err
	}
	return &staff, nil
}

// UpdateStaff met à jour les informations d’un staff
func UpdateStaff(db *gorm.DB, staff *Staff) error {
	return db.Save(staff).Error
}

// DeleteStaff supprime un membre du personnel (soft delete)
func DeleteStaff(db *gorm.DB, id uint64) error {
	return db.Delete(&Staff{}, id).Error
}

// ListStaff retourne la liste du personnel
func ListStaff(db *gorm.DB) ([]Staff, error) {
	var list []Staff
	err := db.Order("last_name ASC").Find(&list).Error
	return list, err
}

// Nombre de staff existant dans la database
func CountStaff(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Staff{}).Count(&count).Error
	return count, err
}
