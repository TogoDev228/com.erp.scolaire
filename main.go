package main

import (
	"school/config"
	"school/controllers"
	"school/models"
	"time"

	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	db := config.InitDB()

	// Crée la table si elle n'existe pas
	db.AutoMigrate(
		&models.User{},
		&models.Student{},
		&models.Staff{},
		&models.Prof{},
		&models.Payment{},
		&models.Parent{},
		&models.Log{},
		&models.Lesson{},
		&models.TransactionHistory{},
		&models.Class{},
		&models.Activity{},
		&models.Item{},
		&models.SchoolYear{},
	)

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"formatDate": func(t interface{}) string {
			if t == nil {
				return "Aucune"
			}

			if tt, ok := t.(time.Time); ok {
				// Vérifie que ce n'est pas une valeur zéro
				if tt.IsZero() {
					return "Aucune"
				}
				return tt.Format("2006-01-02")
			}

			return "Aucune"
		},
	})

	r.Static("/js", "./views/js")

	r.LoadHTMLGlob("views/*.html")

	// Routes auth
	r.GET("/register", controllers.ShowRegister)
	r.POST("/register", func(c *gin.Context) { controllers.Register(c, db) })

	r.GET("/login", controllers.ShowLogin)
	r.POST("/login", func(c *gin.Context) { controllers.Login(c, db) })

	r.GET("/logout", controllers.Logout)

	// Routes protégées
	auth := r.Group("/")
	auth.Use(RequireAuth())

	auth.GET("/dashboard-default", func(c *gin.Context) {
		controllers.ShowDashboard(c, db)
	})

	auth.GET("/dashboard-treasury", func(c *gin.Context) {
		controllers.ShowTreasury(c, db)
	})

	auth.GET("/dashboard-item", func(c *gin.Context) {
		controllers.ShowItem(c, db)
	})

	auth.GET("/dashboard-statistic", func(c *gin.Context) {
		controllers.ShowStatistic(c, db)
	})

	auth.GET("/prof-list", func(c *gin.Context) {
		controllers.ShowProfList(c, db)
	})

	auth.GET("/student-list", func(c *gin.Context) {
		controllers.ShowStudentList(c, db)
	})

	auth.GET("/parent-list", func(c *gin.Context) {
		controllers.ShowParentList(c, db)
	})

	auth.GET("/staff-list", func(c *gin.Context) {
		controllers.ShowStaffList(c, db)
	})

	auth.GET("/activity-list", func(c *gin.Context) {
		controllers.ShowActivityList(c, db)
	})

	auth.GET("/item-list", func(c *gin.Context) {
		controllers.ShowItemList(c, db)
	})

	auth.GET("/setting-class", func(c *gin.Context) {
		controllers.ShowSettingClassroom(c, db)
	})

	auth.GET("/setting-lesson", func(c *gin.Context) {
		controllers.ShowSettingLesson(c, db)
	})

	auth.GET("/setting-log", func(c *gin.Context) {
		controllers.ShowSettingLog(c, db)
	})

	auth.GET("/setting-security", func(c *gin.Context) {
		controllers.ShowSettingSecuriy(c, db)
	})

	auth.GET("/setting-general", func(c *gin.Context) {
		controllers.ShowSettingGeneral(c, db)
	})

	auth.GET("/setting-notification", func(c *gin.Context) {
		controllers.ShowSettingNotification(c, db)
	})

	// POST
	auth.POST("/setting-class", func(c *gin.Context) {
		controllers.AddClass(c, db)
	})

	auth.POST("/setting-lesson", func(c *gin.Context) {
		controllers.AddLesson(c, db)
	})

	auth.POST("/prof-list", func(c *gin.Context) {
		controllers.AddProf(c, db)
	})

	auth.POST("/parent-list", func(c *gin.Context) {
		controllers.AddParent(c, db)
	})

	auth.POST("/staff-list", func(c *gin.Context) {
		controllers.AddStaff(c, db)
	})

	auth.POST("/student-list", func(c *gin.Context) {
		controllers.AddStudent(c, db)
	})

	auth.POST("/activity-list", func(c *gin.Context) {
		controllers.AddActivity(c, db)
	})

	auth.POST("/item-list", func(c *gin.Context) {
		controllers.AddItem(c, db)
	})

	// PUT
	auth.POST("/prof/update/:id", func(c *gin.Context) {
		controllers.UpdateProf(c, db)
	})

	auth.POST("/student/update/:id", func(c *gin.Context) {
		controllers.UpdateStudent(c, db)
	})

	auth.POST("/parent/update/:id", func(c *gin.Context) {
		controllers.UpdateParent(c, db)
	})

	auth.POST("/staff/update/:id", func(c *gin.Context) {
		controllers.UpdateStaff(c, db)
	})

	auth.POST("/activity/update/:id", func(c *gin.Context) {
		controllers.UpdateActivity(c, db)
	})

	auth.POST("/item/update/:id", func(c *gin.Context) {
		controllers.UpdateItem(c, db)
	})

	// DELETE
	auth.POST("/prof/delete/:id", func(c *gin.Context) {
		controllers.DeleteProf(c, db)
	})

	auth.POST("/student/delete/:id", func(c *gin.Context) {
		controllers.DeleteStudent(c, db)
	})

	auth.POST("/parent/delete/:id", func(c *gin.Context) {
		controllers.DeleteParent(c, db)
	})

	auth.POST("/staff/delete/:id", func(c *gin.Context) {
		controllers.DeleteStaff(c, db)
	})

	auth.POST("/activity/delete/:id", func(c *gin.Context) {
		controllers.DeleteActivity(c, db)
	})

	auth.POST("/item/delete/:id", func(c *gin.Context) {
		controllers.DeleteItem(c, db)
	})

	r.Run(":8080")
}

// Middleware
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := c.Cookie("user_id")
		if err != nil || userID == "" {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
