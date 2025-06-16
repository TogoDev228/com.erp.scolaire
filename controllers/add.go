package controllers

import (
	"net/http"
	"strconv"
	"time"

	"school/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddClass(c *gin.Context, db *gorm.DB) {
	title := c.PostForm("title")
	class := &models.Class{Title: title}

	if err := models.CreateClass(db, class); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Class " + title + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Classe " + title + " créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/setting-class")
}

func AddLesson(c *gin.Context, db *gorm.DB) {
	title := c.PostForm("title")
	lesson := &models.Lesson{Title: title}

	if err := models.CreateLesson(db, lesson); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Matière " + title + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Matière " + title + " créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/setting-lesson")
}

func AddParent(c *gin.Context, db *gorm.DB) {
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	sexe := c.PostForm("sexe")
	grade := c.PostForm("grade")
	email := c.PostForm("email")
	position := c.PostForm("job")
	phone := c.PostForm("phone")
	startStr := c.PostForm("start")
	childrenStr := c.PostForm("children")

	// Conversion de start en time.Time
	start, err := time.Parse("2006-01-02", startStr) // format exemple : "2025-06-05"
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	// Conversion de children en int
	children, err := strconv.Atoi(childrenStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Nombre d'enfants invalide : %v", err)
		return
	}

	parent := &models.Parent{FirstName: firstName, LastName: lastName, Sexe: sexe, Grade: grade, Position: position, Email: email, Phone: phone, Start: start, Children: children}

	if err := models.CreateParent(db, parent); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Parent " + firstName + lastName + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Parent " + firstName + lastName + " créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/parent-list")
}

func AddStudent(c *gin.Context, db *gorm.DB) {
	matricul := c.PostForm("matricul")
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	sexe := c.PostForm("sexe")
	grade := c.PostForm("grade")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	startStr := c.PostForm("start")
	statut := c.PostForm("statut")
	typeInscription := c.PostForm("inscription")
	typ := c.PostForm("type")
	matrimonial := c.PostForm("matrimonial")
	schoolYear := c.PostForm("schoolYear")

	// Conversion de start en time.Time
	start, err := time.Parse("2006-01-02", startStr) // format exemple : "2025-06-05"
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	student := &models.Student{Matricul: matricul, FirstName: firstName, LastName: lastName, Sexe: sexe, Grade: grade, Email: email, Phone: phone, Start: start, Statut: statut, TypeInscription: typeInscription, Type: typ, Matrimonial: matrimonial, SchoolYear: schoolYear}

	if err := models.CreateStudent(db, student); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Apprenant " + firstName + lastName + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Apprenant " + firstName + lastName + " créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/student-list")
}

func AddActivity(c *gin.Context, db *gorm.DB) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	typ := c.PostForm("type")
	budget := c.PostForm("budget")
	location := c.PostForm("location")
	startStr := c.PostForm("start")
	endStr := c.PostForm("end")

	// Conversion de start en time.Time
	start, err := time.Parse("2006-01-02", startStr) // format exemple : "2025-06-05"
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide - (start) : %v", err)
		return
	}

	// Conversion de end en time.Time
	end, err := time.Parse("2006-01-02", endStr) // format exemple : "2025-06-05"
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide - (end) : %v", err)
		return
	}

	activity := &models.Activity{Title: title, Description: description, Type: typ, Budget: budget, Location: location, Start: start, End: end}

	if err := models.CreateActivity(db, activity); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Activité " + title + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Activité " + title + " créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/activity-list")
}

func AddProf(c *gin.Context, db *gorm.DB) {
	matricul := c.PostForm("matricul")
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	sexe := c.PostForm("sexe")
	grade := c.PostForm("grade")
	email := c.PostForm("email")
	lesson := c.PostForm("lesson")
	phone := c.PostForm("phone")
	startStr := c.PostForm("start")

	// Conversion de start en time.Time
	start, err := time.Parse("2006-01-02", startStr) // format exemple : "2025-06-05"
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	prof := &models.Prof{Matricul: matricul, FirstName: firstName, LastName: lastName, Lesson: lesson, Sexe: sexe, Grade: grade, Email: email, Phone: phone, Start: start}

	if err := models.CreateProf(db, prof); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Prof " + firstName + lastName + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Prof " + firstName + lastName + " créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/prof-list")
}

func AddStaff(c *gin.Context, db *gorm.DB) {
	matricul := c.PostForm("matricul")
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	sexe := c.PostForm("sexe")
	grade := c.PostForm("grade")
	email := c.PostForm("email")
	position := c.PostForm("position")
	phone := c.PostForm("phone")
	startStr := c.PostForm("start")

	// Conversion de start en time.Time
	start, err := time.Parse("2006-01-02", startStr) // format exemple : "2025-06-05"
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	staff := &models.Staff{Matricul: matricul, FirstName: firstName, LastName: lastName, Position: position, Sexe: sexe, Grade: grade, Email: email, Phone: phone, Start: start}

	if err := models.CreateStaff(db, staff); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Employé(e) " + firstName + lastName + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Employé(e) " + firstName + lastName + " créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/staff-list")
}

func AddItem(c *gin.Context, db *gorm.DB) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	typ := c.PostForm("type")
	value := c.PostForm("value")
	quantityStr := c.PostForm("quantity")
	statut := c.PostForm("statut")
	startStr := c.PostForm("start")

	// Conversion de start en time.Time
	start, err := time.Parse("2006-01-02", startStr) // format exemple : "2025-06-05"
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Quantity invalide : %v", err)
		return
	}

	item := &models.Item{Title: title, Description: description, Type: typ, Value: value, Quantity: quantity, Status: statut, Start: start}

	if err := models.CreateItem(db, item); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Equipement " + title + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Equipement " + title + " créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/item-list")
}
