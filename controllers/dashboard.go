package controllers

import (
	"net/http"
	"school/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowDashboard(c *gin.Context, db *gorm.DB) {
	lesson, err := models.ListLessonsLimit6(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Impossible de charger les leçons.")
		return
	}

	countActivities, err := models.CountActivities(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Impossible d'obtenir le nombre d'activité.")
		return
	}

	countProfs, err := models.CountProfs(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Impossible d'obtenir le nombre de professeurs.")
		return
	}

	countStudent, err := models.CountStudent(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Impossible d'obtenir le nombre d'étudiants.")
		return
	}

	c.HTML(http.StatusOK, "dashboard-default.html", gin.H{
		"lesson":          lesson,
		"countActivities": countActivities,
		"countProfs":      countProfs,
		"countStudent":    countStudent,
	})
}

func ShowTreasury(c *gin.Context, db *gorm.DB) {
	prof, err := models.ListProfsLimit6(db)

	if err != nil {
		c.String(http.StatusInternalServerError, "message, impossible de charger la liste des professeurs.")
		return
	}

	c.HTML(http.StatusOK, "dashboard-treasury.html", gin.H{
		"prof": prof,
	})
}

func ShowItem(c *gin.Context, db *gorm.DB) {
	item, err := models.ListItemsLimit6(db)

	if err != nil {
		c.String(http.StatusInternalServerError, "message, impossible de charger la lise des équipements.")
		return
	}

	countItemsIT, err := models.CountItemsIT(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Impossible d'obtenir le nombre d'équipement.")
		return
	}

	countItemsEducational, err := models.CountItemsEducational(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Impossible d'obtenir le nombre de log.")
		return
	}

	countItemsAdministratifs, err := models.CountItemsAdministratifs(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Impossible d'obtenir le nombre de log.")
		return
	}

	c.HTML(http.StatusOK, "dashboard-item.html", gin.H{
		"item":                     item,
		"countItemsIT":             countItemsIT,
		"countItemsEducational":    countItemsEducational,
		"countItemsAdministratifs": countItemsAdministratifs,
	})
}

func ShowStatistic(c *gin.Context, db *gorm.DB) {

	countItems, err := models.CountItems(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Impossible d'obtenir le nombre d'équipement.")
		return
	}

	countLogs, err := models.CountLogs(db)
	if err != nil {
		c.String(http.StatusInternalServerError, "Impossible d'obtenir le nombre de log.")
		return
	}

	c.HTML(http.StatusOK, "dashboard-statistic.html", gin.H{
		"countLogs":  countLogs,
		"countItems": countItems,
	})
}
