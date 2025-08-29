package models

import (
	"time"

	"gorm.io/gorm"
)

type Attendance struct {
	ID         uint64      `gorm:"primaryKey"`
	UserID     uint64       `gorm:"not null"`
	UserType   string    `gorm:"size:100;not null"`
	Type       string    `gorm:"size:50;not null"`
	Date       time.Time `gorm:"not null"`
	SchoolYear string    `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// Ajoute un nouvau conger pour un utilisateur donnée
func CreateAttendance(db *gorm.DB, attendance *Attendance) error {
	return db.Create(attendance).Error
}

// Récupère un conger via son ID
func GetAttendanceByID(db *gorm.DB, id uint) (*Attendance, error) {
	var attendance Attendance
	err := db.First(&attendance, id).Error
	if err != nil {
		return nil, err
	}
	return &attendance, nil
}

// Modifie un conger existant
func UpdateAttendance(db *gorm.DB, attendance *Attendance) error {
	return db.Save(attendance).Error
}

// Supprime un conger (soft delete si activé)
func DeleteAttendance(db *gorm.DB, id uint) error {
	return db.Delete(&Attendance{}, id).Error
}

// Retourne tous les conger (optionnel : avec pagination ou filtre par année)
func ListAttendances(db *gorm.DB) ([]Attendance, error) {
	var attendances []Attendance
	err := db.Order("last_name ASC").Find(&attendances).Error
	return attendances, err
}

// Nombre de conger existant dans la database
func CountAttendance(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Attendance{}).Count(&count).Error
	return count, err
}
