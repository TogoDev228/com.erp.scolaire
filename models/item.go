package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint       `gorm:"primaryKey"`
	Title       string     `gorm:"size:100;not null"` // Nom de l’objet
	Description string     `gorm:"type:text"`         // Détails optionnels
	Type        string     `gorm:"size:50;not null"`  // Ex: "Equipements sportifs", "Equipements de sécurité", "Equipements technologiques"
	Value       string     `gorm:"size:50"`           // Prix ou référence
	Quantity    int        // Quantité disponible
	Status      string     `gorm:"size:50"`  // Ex: "Bon", "Réparé", "HS"
	Start       time.Time  `gorm:"not null"` // Date d’ajout/mise en service
	Repair      *time.Time // Date de dernière réparation
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"` // Pour soft delete
}

// CreateItem ajoute un nouvel item
func CreateItem(db *gorm.DB, item *Item) error {
	return db.Create(item).Error
}

// GetItemByID récupère un item par ID
func GetItemByID(db *gorm.DB, id uint) (*Item, error) {
	var item Item
	err := db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// ListItems retourne tous les items
func ListItems(db *gorm.DB) ([]Item, error) {
	var items []Item
	err := db.Order("title ASC").Find(&items).Error
	return items, err
}

// ListItems retourne tous les items avec un limit de 6
func ListItemsLimit6(db *gorm.DB) ([]Item, error) {
	var items []Item
	err := db.Order("title ASC").Limit(6).Find(&items).Error
	return items, err
}

// UpdateItem met à jour un item
func UpdateItem(db *gorm.DB, item *Item) error {
	return db.Save(item).Error
}

// DeleteItem supprime un item (soft delete)
func DeleteItem(db *gorm.DB, id uint) error {
	return db.Delete(&Item{}, id).Error
}

// Nombre de prof existant dans la database
func CountItems(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Item{}).Count(&count).Error
	return count, err
}

// Nombre de prof existant dans la database
func CountItemsIT(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Item{}).Where("type = ?", "Equipements technologiques").Count(&count).Error
	return count, err
}

// Nombre de prof existant dans la database
func CountItemsEducational(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Item{}).Where("type = ?", "Equipements pédagogiques").Count(&count).Error
	return count, err
}

// Nombre de prof existant dans la database
func CountItemsAdministratifs(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Item{}).Where("type = ?", "Equipements administratifs").Count(&count).Error
	return count, err
}
