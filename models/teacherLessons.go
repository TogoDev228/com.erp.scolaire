package models

import (
	"time"

	"gorm.io/gorm"
)

type TeacherLesson struct {
	ID         uint   `gorm:"primaryKey"`
	TeacherID  int    `gorm:"not null"`
	Lesson     string `gorm:"not null"`
	SchoolYear string `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// Crée une nouvelle relation entre un professeur et une lesson
func CreateTeacherLesson(db *gorm.DB, teacherLesson *TeacherLesson) error {
	return db.Create(teacherLesson).Error
}

// Récupère une relation entre un professeur et une lesson
func GetTeacherLessonByID(db *gorm.DB, id uint) (*TeacherLesson, error) {
	var teacherLesson TeacherLesson
	err := db.First(&teacherLesson, id).Error
	if err != nil {
		return nil, err
	}
	return &teacherLesson, nil
}

// Modifie une relation entre un professeur et une lesson
func UpdateTeacherLesson(db *gorm.DB, teacherLesson *TeacherLesson) error {
	return db.Save(teacherLesson).Error
}

// Supprime une relation entre un professeur et une lesson
func DeleteTeacherLesson(db *gorm.DB, id uint) error {
	return db.Delete(&TeacherLesson{}, id).Error
}

// Liste toutes les relation entre un professeur et des lessons
func ListTeacherLessons(db *gorm.DB) ([]TeacherLesson, error) {
	var teacherLessons []TeacherLesson
	err := db.Order("start DESC").Find(&teacherLessons).Error
	return teacherLessons, err
}
