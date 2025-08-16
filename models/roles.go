package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"size:100;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// Ajoute une nouvelle Role
func CreateRole(db *gorm.DB, role *Role) error {
	return db.Create(role).Error
}

// Récupère une Role par ID
func GetRoleByID(db *gorm.DB, id uint) (*Role, error) {
	var role Role
	err := db.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

// Retourne toutes les Roles (ordre alphabétique)
func ListRoles(db *gorm.DB) ([]Role, error) {
	var role []Role
	err := db.Order("title ASC").Find(&role).Error
	return role, err
}

// Met à jour un Role
func UpdateRole(db *gorm.DB, role *Role) error {
	return db.Save(role).Error
}

// Supprime une Role (soft delete)
func DeleteRole(db *gorm.DB, id uint) error {
	return db.Delete(&Role{}, id).Error
}

// Nombre de Role existant dans la database
func CountRole(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Role{}).Count(&count).Error
	return count, err
}
