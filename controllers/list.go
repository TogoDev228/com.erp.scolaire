package controllers

import (
	"net/http"
	"school/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowProfList(c *gin.Context, db *gorm.DB) {
	prof, err := models.ListProfs(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "message, impossible de charger la liste des professeurs.")
		return
	}

	c.HTML(http.StatusOK, "list-prof.html", gin.H{
		"prof": prof,
	})
}

func ShowStudentList(c *gin.Context, db *gorm.DB) {
	student, err := models.ListStudents(db)

	if err != nil {
		c.String(http.StatusInternalServerError, "message, impossible de charger la liste des apprenants.")
		return
	}

	class, err := models.ListClasses(db)

	if err != nil {
		c.String(http.StatusInternalServerError, "message, impossible de charger les classes.")
		return
	}

	c.HTML(http.StatusOK, "list-student.html", gin.H{
		"student": student,
		"class":   class,
	})
}

func ShowParentList(c *gin.Context, db *gorm.DB) {
	parent, err := models.ListParents(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "message, impossible de charger la liste des parents.")
		return
	}

	c.HTML(http.StatusOK, "list-parent.html", gin.H{
		"parent": parent,
	})
}

func ShowStaffList(c *gin.Context, db *gorm.DB) {
	staff, err := models.ListStaff(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "message, impossible de charger la liste des employés.")
		return
	}

	c.HTML(http.StatusOK, "list-staff.html", gin.H{
		"staff": staff,
	})
}

func ShowActivityList(c *gin.Context, db *gorm.DB) {
	activity, err := models.ListActivities(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "message, impossible de charger la liste des activités.")
		return
	}

	c.HTML(http.StatusOK, "list-activity.html", gin.H{
		"activity": activity,
	})
}

func ShowItemList(c *gin.Context, db *gorm.DB) {
	item, err := models.ListItems(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "message, impossible de charger la liste des équipements.")
		return
	}

	c.HTML(http.StatusOK, "list-item.html", gin.H{
		"item": item,
	})
}
