package models

import (
	"time"

	"gorm.io/gorm"
)

type Prof struct {
	ID        uint      `gorm:"primaryKey"`
	Matricul  string    `gorm:"size:50;unique;not null"`
	FirstName string    `gorm:"size:100;not null"`
	LastName  string    `gorm:"size:100;not null"`
	Sexe      string    `gorm:"size:10"`
	Email     string    `gorm:"size:100;unique"`
	Grade     string    `gorm:"size:50"` // ex: "Agrégé", "Contractuel"
	Phone     string    `gorm:"size:20"`
	Start     time.Time `gorm:"not null"` // Date d’entrée
	End       *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete
}

// CreateProf ajoute un professeur
func CreateProf(db *gorm.DB, prof *Prof) error {
	return db.Create(prof).Error
}

// GetProfByID récupère un prof par ID
func GetProfByID(db *gorm.DB, id uint) (*Prof, error) {
	var prof Prof
	err := db.First(&prof, id).Error
	if err != nil {
		return nil, err
	}
	return &prof, nil
}

// GetProfByMatricul récupère un prof via son matricule
func GetProfByMatricul(db *gorm.DB, matricul string) (*Prof, error) {
	var prof Prof
	err := db.Where("matricul = ?", matricul).First(&prof).Error
	if err != nil {
		return nil, err
	}
	return &prof, nil
}

// UpdateProf met à jour un prof
func UpdateProf(db *gorm.DB, prof *Prof) error {
	return db.Save(prof).Error
}

// DeleteProf supprime un prof (soft delete)
func DeleteProf(db *gorm.DB, id uint) error {
	return db.Delete(&Prof{}, id).Error
}

// ListProfs retourne tous les professeurs
func ListProfs(db *gorm.DB) ([]Prof, error) {
	var list []Prof
	err := db.Order("last_name ASC").Find(&list).Error
	return list, err
}

// ListProfs retourne tous les professeurs avec une limite de 6
func ListProfsLimit6(db *gorm.DB) ([]Prof, error) {
	var list []Prof
	err := db.Order("last_name ASC").Limit(6).Find(&list).Error
	return list, err
}

// Nombre de prof existant dans la database
func CountProfs(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Prof{}).Count(&count).Error
	return count, err
}
