package controllers

import (
	"net/http"

	"school/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ShowRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func Register(c *gin.Context, db *gorm.DB) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := &models.User{Username: name, Email: email, Password: password}
	if err := models.CreateUser(db, user); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/login")
}

func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context, db *gorm.DB) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := models.FindUserByEmail(db, email)
	if err != nil || !models.CheckPasswordHash(password, user.Password) {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "Email ou mot de passe incorrect"})
		return
	}

	// Simuler une session simple avec un cookie
	c.SetCookie("user_id", string(rune(user.ID)), 10000, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/dashboard-default")
}

func Logout(c *gin.Context) {
	c.SetCookie("user_id", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
}

func ShowAllUsers(c *gin.Context, db *gorm.DB) {
	var users []models.User

	if err := db.Find(&users).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "users.html", gin.H{
			"Users": []models.User{},
		})
		return
	}

	c.HTML(http.StatusOK, "users.html", gin.H{
		"Users": users,
	})
}
