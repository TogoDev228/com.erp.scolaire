package controllers

import (
	"net/http"
	"school/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowSettingClassroom(c *gin.Context, db *gorm.DB) {
	class, err := models.ListClasses(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "message, impossible de charger les classes.")
		return
	}

	c.HTML(http.StatusOK, "setting-class.html", gin.H{
		"class": class,
	})
}

func ShowSettingLesson(c *gin.Context, db *gorm.DB) {
	lesson, err := models.ListLessons(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "message, impossible de charger les matières.")
		return
	}

	c.HTML(http.StatusOK, "setting-lesson.html", gin.H{
		"lesson": lesson,
	})
}

func ShowSettingLog(c *gin.Context, db *gorm.DB) {
	logs, err := models.ListLogs(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "message, impossible de charger les logs.")
		return
	}

	c.HTML(http.StatusOK, "setting-log.html", gin.H{
		"logs": logs,
	})
}

func ShowSettingSecuriy(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "setting-security.html", nil)
}

func ShowSettingGeneral(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "setting-general.html", nil)
}

func ShowSettingNotification(c *gin.Context, db *gorm.DB) {
	c.HTML(http.StatusOK, "setting-notification.html", nil)
}
