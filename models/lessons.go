package models

import (
	"time"

	"gorm.io/gorm"
)

type Lesson struct {
	ID        uint64   `gorm:"primaryKey"`
	Title     string `gorm:"size:100;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// CreateLesson ajoute une leçon
func CreateLesson(db *gorm.DB, lesson *Lesson) error {
	return db.Create(lesson).Error
}

// GetLessonByID récupère une leçon par son ID
func GetLessonByID(db *gorm.DB, id uint64) (*Lesson, error) {
	var lesson Lesson
	err := db.First(&lesson, id).Error
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

// ListLessons liste toutes les leçons
func ListLessons(db *gorm.DB) ([]Lesson, error) {
	var lessons []Lesson
	err := db.Order("title ASC").Find(&lessons).Error
	return lessons, err
}

// ListLessons liste toutes les leçons avec une limite de 6
func ListLessonsLimit6(db *gorm.DB) ([]Lesson, error) {
	var lessons []Lesson
	err := db.Order("title ASC").Limit(6).Find(&lessons).Error
	return lessons, err
}

// UpdateLesson met à jour une leçon
func UpdateLesson(db *gorm.DB, lesson *Lesson) error {
	return db.Save(lesson).Error
}

// DeleteLesson supprime une leçon (soft delete)
func DeleteLesson(db *gorm.DB, id uint64) error {
	return db.Delete(&Lesson{}, id).Error
}

// Nombre de Lesson existant dans la database
func CountLessons(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Lesson{}).Count(&count).Error
	return count, err
}
