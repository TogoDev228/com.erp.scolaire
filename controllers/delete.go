package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"school/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteClass(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)

	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteClass(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("Classe ID %d n'a pas pu être supprimée", id)}
		_ = models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("Classe ID %d supprimée avec succès", id)}
	_ = models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/setting-class")
}

func DeleteLesson(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteLesson(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: fmt.Sprintf("Leçon ID %d non supprimée", id)})
		return
	}

	models.CreateLog(db, &models.Log{Type: "DELETE", Message: fmt.Sprintf("Leçon ID %d supprimée", id)})
	c.Redirect(http.StatusSeeOther, "/setting-lesson")
}

func DeleteParent(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteParent(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: fmt.Sprintf("Parent ID %d non supprimé", id)})
		return
	}

	models.CreateLog(db, &models.Log{Type: "DELETE", Message: fmt.Sprintf("Parent ID %d supprimé", id)})
	c.Redirect(http.StatusSeeOther, "/parent-list")
}

func DeleteStudent(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteStudent(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: fmt.Sprintf("Apprenant ID %d non supprimé", id)})
		return
	}

	models.CreateLog(db, &models.Log{Type: "DELETE", Message: fmt.Sprintf("Apprenant ID %d supprimé", id)})
	c.Redirect(http.StatusSeeOther, "/student-list")
}

func DeleteActivity(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteActivity(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: fmt.Sprintf("Activité ID %d non supprimée", id)})
		return
	}

	models.CreateLog(db, &models.Log{Type: "DELETE", Message: fmt.Sprintf("Activité ID %d supprimée", id)})
	c.Redirect(http.StatusSeeOther, "/activity-list")
}

func DeleteProf(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteProf(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: fmt.Sprintf("Prof ID %d non supprimé", id)})
		return
	}

	models.CreateLog(db, &models.Log{Type: "DELETE", Message: fmt.Sprintf("Prof ID %d supprimé", id)})
	c.Redirect(http.StatusSeeOther, "/prof-list")
}

func DeleteStaff(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteStaff(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: fmt.Sprintf("Staff ID %d non supprimé", id)})
		return
	}

	models.CreateLog(db, &models.Log{Type: "DELETE", Message: fmt.Sprintf("Staff ID %d supprimé", id)})
	c.Redirect(http.StatusSeeOther, "/staff-list")
}

func DeleteItem(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteItem(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: fmt.Sprintf("Item ID %d non supprimé", id)})
		return
	}

	models.CreateLog(db, &models.Log{Type: "DELETE", Message: fmt.Sprintf("Item ID %d supprimé", id)})
	c.Redirect(http.StatusSeeOther, "/item-list")
}

func getIDParam(c *gin.Context) (uint, error) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(idUint), nil
}
