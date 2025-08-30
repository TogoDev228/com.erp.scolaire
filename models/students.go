package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID              uint64    `gorm:"primaryKey"`
	Matricul        string    `gorm:"size:50;unique;not null"` // Numéro matricule unique
	FirstName       string    `gorm:"size:100;not null"`
	LastName        string    `gorm:"size:100;not null"`
	Sexe            string    `gorm:"size:10"`          // "Masculin" / "Feminin" ou "Male" / "Female"
	Grade           string    `gorm:"size:50"`          // Classe / niveau
	Phone           string    `gorm:"size:20"`          // Optionnel
	Email           string    `gorm:"size:20"`          // Optionnel
	Start           time.Time `gorm:"not null"`         // Date d'inscription
	Statut          string    `gorm:"not null"`         // Inscript, diplomé, Exclu,...
	TypeInscription string    `gorm:"not null"`         // Initial, Réinscript,...
	Type            string    `gorm:"not null"`         // Régulier, Boursier, Handicapé ...
	Matrimonial     string    `gorm:"not null"`         // Célibataire, Marié, Veuf, Divorcé
	SchoolYear      string    `gorm:"size:20;not null"` // ex: "2024-2025"
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"` // Pour soft delete (optionnel)
}

// CreateStudent ajoute un nouvel étudiant
func CreateStudent(db *gorm.DB, student *Student) error {
	return db.Create(student).Error
}

// GetStudentByID récupère un étudiant via son ID
func GetStudentByID(db *gorm.DB, id uint64) (*Student, error) {
	var student Student
	err := db.First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

// UpdateStudent modifie un étudiant existant
func UpdateStudent(db *gorm.DB, student *Student) error {
	return db.Save(student).Error
}

// DeleteStudent supprime un étudiant (soft delete si activé)
func DeleteStudent(db *gorm.DB, id uint64) error {
	return db.Delete(&Student{}, id).Error
}

// GetStudentByMatricul récupère un étudiant par matricule
func GetStudentByMatricul(db *gorm.DB, matricul string) (*Student, error) {
	var student Student
	err := db.Where("matricul = ?", matricul).First(&student).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

// ListStudents retourne tous les étudiants (optionnel : avec pagination ou filtre par année)
func ListStudents(db *gorm.DB) ([]Student, error) {
	var students []Student
	err := db.Order("last_name ASC").Find(&students).Error
	return students, err
}

// Nombre de prof existant dans la database
func CountStudent(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Student{}).Count(&count).Error
	return count, err
}
