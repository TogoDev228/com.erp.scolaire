package controllers

import (
	"net/http"
	"strconv"
	"time"

	"school/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateClass(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id") // On récupère l'ID dans l'URL, ex: /class/update/:id
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	title := c.PostForm("title")

	class := &models.Class{}
	if err := db.First(class, id).Error; err != nil {
		c.String(http.StatusNotFound, "Classe non trouvée : %v", err)
		return
	}

	class.Title = title

	if err := db.Save(class).Error; err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la mise à jour : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Classe " + title + " n'a pas pu être mise à jour"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
		}
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "Classe " + title + " mise à jour avec succès"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/setting-class")
}

func UpdateLesson(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	title := c.PostForm("title")

	lesson := &models.Lesson{}
	if err := db.First(lesson, id).Error; err != nil {
		c.String(http.StatusNotFound, "Matière non trouvée : %v", err)
		return
	}

	lesson.Title = title

	if err := db.Save(lesson).Error; err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la mise à jour : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Matière " + title + " n'a pas pu être mise à jour"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
		}
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "Matière " + title + " mise à jour avec succès"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/setting-lesson")
}

func UpdateParent(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	sexe := c.PostForm("sexe")
	grade := c.PostForm("grade")
	email := c.PostForm("email")
	position := c.PostForm("job")
	phone := c.PostForm("phone")
	startStr := c.PostForm("start")
	childrenStr := c.PostForm("children")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	children, err := strconv.Atoi(childrenStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Nombre d'enfants invalide : %v", err)
		return
	}

	parent := &models.Parent{}
	if err := db.First(parent, id).Error; err != nil {
		c.String(http.StatusNotFound, "Parent non trouvé : %v", err)
		return
	}

	parent.FirstName = firstName
	parent.LastName = lastName
	parent.Sexe = sexe
	parent.Grade = grade
	parent.Position = position
	parent.Email = email
	parent.Phone = phone
	parent.Start = start
	parent.Children = children

	if err := db.Save(parent).Error; err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la mise à jour : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Parent " + firstName + lastName + " n'a pas pu être mis à jour"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
		}
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "Parent " + firstName + lastName + " mis à jour avec succès"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/parent-list")
}

func UpdateStudent(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

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

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	student := &models.Student{}
	if err := db.First(student, id).Error; err != nil {
		c.String(http.StatusNotFound, "Apprenant non trouvé : %v", err)
		return
	}

	student.Matricul = matricul
	student.FirstName = firstName
	student.LastName = lastName
	student.Sexe = sexe
	student.Grade = grade
	student.Email = email
	student.Phone = phone
	student.Start = start
	student.Statut = statut
	student.TypeInscription = typeInscription
	student.Type = typ
	student.Matrimonial = matrimonial
	student.SchoolYear = schoolYear

	if err := db.Save(student).Error; err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la mise à jour : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Apprenant " + firstName + lastName + " n'a pas pu être mis à jour"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
		}
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "Apprenant " + firstName + lastName + " mis à jour avec succès"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/student-list")
}

func UpdateActivity(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("description")
	typ := c.PostForm("type")
	budget := c.PostForm("budget")
	location := c.PostForm("location")
	startStr := c.PostForm("start")
	endStr := c.PostForm("end")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide - (start) : %v", err)
		return
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide - (end) : %v", err)
		return
	}

	activity := &models.Activity{}
	if err := db.First(activity, id).Error; err != nil {
		c.String(http.StatusNotFound, "Activité non trouvée : %v", err)
		return
	}

	activity.Title = title
	activity.Description = description
	activity.Type = typ
	activity.Budget = budget
	activity.Location = location
	activity.Start = start
	activity.End = end

	if err := db.Save(activity).Error; err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la mise à jour : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Activité " + title + " n'a pas pu être mise à jour"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
		}
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "Activité " + title + " mise à jour avec succès"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/activity-list")
}

func UpdateProf(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	matricul := c.PostForm("matricul")
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	sexe := c.PostForm("sexe")
	grade := c.PostForm("grade")
	email := c.PostForm("email")
	lesson := c.PostForm("lesson")
	phone := c.PostForm("phone")
	startStr := c.PostForm("start")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	prof := &models.Prof{}
	if err := db.First(prof, id).Error; err != nil {
		c.String(http.StatusNotFound, "Prof non trouvé : %v", err)
		return
	}

	prof.Matricul = matricul
	prof.FirstName = firstName
	prof.LastName = lastName
	prof.Sexe = sexe
	prof.Grade = grade
	prof.Email = email
	prof.Lesson = lesson
	prof.Phone = phone
	prof.Start = start

	if err := db.Save(prof).Error; err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la mise à jour : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Prof " + firstName + lastName + " n'a pas pu être mis à jour"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
		}
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "Prof " + firstName + lastName + " mis à jour avec succès"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/prof-list")
}

func UpdateStaff(c *gin.Context, db *gorm.DB) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	matricul := c.PostForm("matricul")
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	sexe := c.PostForm("sexe")
	grade := c.PostForm("grade")
	email := c.PostForm("email")
	position := c.PostForm("position")
	phone := c.PostForm("phone")
	startStr := c.PostForm("start")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	staff := &models.Staff{}
	if err := db.First(staff, id).Error; err != nil {
		c.String(http.StatusNotFound, "Employé(e) non trouvé(e) : %v", err)
		return
	}

	staff.Matricul = matricul
	staff.FirstName = firstName
	staff.LastName = lastName
	staff.Position = position
	staff.Sexe = sexe
	staff.Grade = grade
	staff.Email = email
	staff.Phone = phone
	staff.Start = start

	if err := db.Save(staff).Error; err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la mise à jour : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Employé(e) " + firstName + lastName + " n'a pas pu être mis à jour"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
		}
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "Employé(e) " + firstName + lastName + " mis à jour avec succès"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/staff-list")
}

func UpdateItem(c *gin.Context, db *gorm.DB) {
	// Récupérer l'ID depuis l'URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("description")
	typ := c.PostForm("type")
	value := c.PostForm("value")
	quantityStr := c.PostForm("quantity")
	statut := c.PostForm("statut")
	startStr := c.PostForm("start")

	// Conversion de start en time.Time
	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Quantité invalide : %v", err)
		return
	}

	// Récupérer l'item existant depuis la BDD
	item, err := models.GetItemByID(db, uint(id))
	if err != nil {
		c.String(http.StatusNotFound, "Equipement non trouvé : %v", err)
		return
	}

	// Mettre à jour les champs
	item.Title = title
	item.Description = description
	item.Type = typ
	item.Value = value
	item.Quantity = quantity
	item.Status = statut
	item.Start = start

	// Enregistrer la mise à jour en BDD
	if err := models.UpdateItem(db, item); err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la mise à jour : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Equipement " + title + " n'a pas pu être mis à jour"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur lors de la création du log : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "Equipement " + title + " mis à jour avec succès"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la création du log : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/item-list")
}
