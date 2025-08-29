package models

import (
	"time"

	"gorm.io/gorm"
)

type Exam struct {
	ID         uint64   `gorm:"primaryKey"`
	Title      string `gorm:"not null"`
	Type       string `gorm:"not null"`
	Status     string `gorm:"not null"`
	SchoolYear string `gorm:"not null"`
	Start      time.Time
	End        time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// Crée un examen
func CreateExam(db *gorm.DB, exam *Exam) error {
	return db.Create(exam).Error
}

// Récupère un examen par son ID
func GetExamByID(db *gorm.DB, id uint) (*Exam, error) {
	var exam Exam
	err := db.First(&exam, id).Error
	if err != nil {
		return nil, err
	}
	return &exam, nil
}

// Modifie un examen
func UpdateExam(db *gorm.DB, exam *Exam) error {
	return db.Save(exam).Error
}

// Supprime un examen
func DeleteExam(db *gorm.DB, id uint) error {
	return db.Delete(&Exam{}, id).Error
}

// Liste toutes les examens dans la base de données
func ListExams(db *gorm.DB) ([]Exam, error) {
	var exams []Exam
	err := db.Order("start DESC").Find(&exams).Error
	return exams, err
}
