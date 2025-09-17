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

func DeletePayment(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeletePayment(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("Le paiement #%d n'a pas pu être supprimé", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("Le paiement #%d a été supprimé avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/payment-list")
}

func DeleteTeacherLesson(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteTeacherLesson(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("La relation prof-leçon #%d n'a pas pu être supprimée", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("La relation prof-leçon #%d a été supprimée avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/teacher-lesson-list")
}

func DeleteAttendance(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteAttendance(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("La présence #%d n'a pas pu être supprimée", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("La présence #%d a été supprimée avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/attendance-list")
}

func DeleteRole(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteRole(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("Le rôle #%d n'a pas pu être supprimé", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("Le rôle #%d a été supprimé avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/setting-role")
}

func DeleteExpense(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteExpense(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("La dépense #%d n'a pas pu être supprimée", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("La dépense #%d a été supprimée avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/expense-list")
}

func DeleteTask(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteTask(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("La tâche #%d n'a pas pu être supprimée", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("La tâche #%d a été supprimée avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/task-list")
}

func DeleteExam(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteExam(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("L'examen #%d n'a pas pu être supprimé", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("L'examen #%d a été supprimé avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/exam-list")
}

func DeleteRemediation(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteRemediation(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("La remédiation #%d n'a pas pu être supprimée", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("La remédiation #%d a été supprimée avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/remediation-list")
}

func DeleteSynchronization(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteSynchronization(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("La synchronisation #%d n'a pas pu être supprimée", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("La synchronisation #%d a été supprimée avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/synchronization-list")
}

func DeletePlanning(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeletePlanning(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("Le planning #%d n'a pas pu être supprimé", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("Le planning #%d a été supprimé avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/planning-list")
}

func DeletePlanningSlot(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeletePlanningSlot(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("La cellule de planning #%d n'a pas pu être supprimée", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("La cellule de planning #%d a été supprimée avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/planning-slot-list")
}

func DeleteSchoolYear(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteSchoolYear(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("L'année scolaire #%d n'a pas pu être supprimée", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("L'année scolaire #%d a été supprimée avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/setting-school-year")
}

func DeleteNotification(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteNotification(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("La notification #%d n'a pas pu être supprimée", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("La notification #%d a été supprimée avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/notification-list")
}

func DeleteStudentClass(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteStudentClass(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("L'association élève-classe #%d n'a pas pu être supprimée", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("L'association élève-classe #%d a été supprimée avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/student-class-list")
}

func DeleteRemuneration(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteRemuneration(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("La rémunération #%d n'a pas pu être supprimée", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("La rémunération #%d a été supprimée avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/remuneration-list")
}

func DeleteLeave(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	if err := models.DeleteLeave(db, id); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		log := &models.Log{Type: "ERROR", Message: fmt.Sprintf("Le congé #%d n'a pas pu être supprimé", id)}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "DELETE", Message: fmt.Sprintf("Le congé #%d a été supprimé avec succès", id)}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/leave-list")
}

func getIDParam(c *gin.Context) (uint64, error) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(idUint), nil
}
