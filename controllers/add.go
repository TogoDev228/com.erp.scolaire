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
	phone := c.PostForm("phone")
	startStr := c.PostForm("start")

	// Conversion de start en time.Time
	start, err := time.Parse("2006-01-02", startStr) // format exemple : "2025-06-05"
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	prof := &models.Prof{Matricul: matricul, FirstName: firstName, LastName: lastName, Sexe: sexe, Grade: grade, Email: email, Phone: phone, Start: start}

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
	valueStr := c.PostForm("value")
	quantityStr := c.PostForm("quantity")
	statut := c.PostForm("statut")
	startStr := c.PostForm("start")

	value, err := strconv.Atoi(valueStr)

	if err != nil {
		c.String(http.StatusBadRequest, "Value invalide : %v", err)
		return
	}

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

func AddPayment(c *gin.Context, db *gorm.DB) {
	studentIDStr := c.PostForm("studentID")
	priceStr := c.PostForm("price")
	typ := c.PostForm("type")

	studentID, err := strconv.ParseUint(studentIDStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "studentID invalide : %v", err)
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "price invalide : %v", err)
		return
	}

	payment := &models.Payment{StudentID: studentID, Price: price, Type: typ}

	if err := models.CreatePayment(db, payment); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Payement de l'apprenant #" + studentIDStr + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Payement de l'apprenant #" + studentIDStr + " créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/payment-list")
}

func AddTeacherLesson(c *gin.Context, db *gorm.DB) {
	teacherIDStr := c.PostForm("studentID")
	lesson := c.PostForm("price")
	schoolYear := c.PostForm("type")

	teacherID, err := strconv.ParseUint(teacherIDStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "teacherID invalide : %v", err)
		return
	}

	teacherLesson := &models.TeacherLesson{TeacherID: teacherID, Lesson: lesson, SchoolYear: schoolYear}

	if err := models.CreateTeacherLesson(db, teacherLesson); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "La relation entre le professeur @" + teacherIDStr + " et la leçon " + lesson + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "La relation entre le professeur @" + teacherIDStr + " et la leçon " + lesson + " à était créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/teacher-lesson-list")
}

func AddAttendance(c *gin.Context, db *gorm.DB) {
	userIDStr := c.PostForm("studentID")
	userType := c.PostForm("price")
	typ := c.PostForm("price")
	dateStr := c.PostForm("price")
	schoolYear := c.PostForm("type")

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "userID invalide : %v", err)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.String(http.StatusBadRequest, "date invalide : %v", err)
		return
	}

	attendance := &models.Attendance{UserID: userID, UserType: userType, Type: typ, Date: date, SchoolYear: schoolYear}

	if err := models.CreateAttendance(db, attendance); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "La présence de l'utilisateur @" + userIDStr + " de " + userType + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "La présence de l'utilisateur @" + userIDStr + " de " + userType + " à était créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/attendance-list")
}

func AddRole(c *gin.Context, db *gorm.DB) {

	title := c.PostForm("title")

	role := &models.Role{Title: title}

	if err := models.CreateRole(db, role); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Le role " + title + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Le role " + title + " à était créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/setting-role")
}

func AddExpense(c *gin.Context, db *gorm.DB) {

	title := c.PostForm("title")
	description := c.PostForm("title")
	typ := c.PostForm("title")
	status := c.PostForm("title")
	schoolYear := c.PostForm("title")
	creator := c.PostForm("title")
	startStr := c.PostForm("title")
	endStr := c.PostForm("title")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de debut invalide : %v", err)
		return
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de fin invalide : %v", err)
		return
	}

	role := &models.Expense{Title: title, Description: description, Type: typ, Status: status, SchoolYear: schoolYear, Creator: creator, Start: start, End: end}

	if err := models.CreateExpense(db, role); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Le role " + title + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Le role " + title + " à était créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/role-list")
}

func AddTask(c *gin.Context, db *gorm.DB) {

	title := c.PostForm("title")
	description := c.PostForm("title")
	typ := c.PostForm("title")
	status := c.PostForm("title")
	schoolYear := c.PostForm("title")
	creator := c.PostForm("title")
	startStr := c.PostForm("title")
	endStr := c.PostForm("title")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de debut invalide : %v", err)
		return
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de fin invalide : %v", err)
		return
	}

	task := &models.Task{Title: title, Description: description, Type: typ, Status: status, SchoolYear: schoolYear, Creator: creator, Start: start, End: end}

	if err := models.CreateTask(db, task); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "La tâche " + title + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Le tâche " + title + " à était créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/task-list")
}

func AddExam(c *gin.Context, db *gorm.DB) {

	title := c.PostForm("title")
	typ := c.PostForm("title")
	Status := c.PostForm("title")
	schoolYear := c.PostForm("title")
	startStr := c.PostForm("title")
	endStr := c.PostForm("title")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de debut invalide : %v", err)
		return
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de fin invalide : %v", err)
		return
	}

	exam := &models.Exam{Title: title, Type: typ, Status: Status, SchoolYear: schoolYear, Start: start, End: end}

	if err := models.CreateExam(db, exam); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "L'examen' " + title + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "L'examen' " + title + " à était créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/exam-list")
}

func AddRemediation(c *gin.Context, db *gorm.DB) {

	title := c.PostForm("title")
	description := c.PostForm("title")
	studentIDStr := c.PostForm("title")
	teacherIDStr := c.PostForm("title")
	Status := c.PostForm("title")
	startStr := c.PostForm("title")
	endStr := c.PostForm("title")


	studentID, err := strconv.ParseUint(studentIDStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "studentID invalide : %v", err)
		return
	}

	teacherID, err := strconv.ParseUint(teacherIDStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "teacherID invalide : %v", err)
		return
	}

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de debut invalide : %v", err)
		return
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de fin invalide : %v", err)
		return
	}

	remediation := &models.Remediation{Title: title, Description: description, StudentID: studentID, TeacherID: teacherID, Status: Status, Start: start, End: end	}

	if err := models.CreateRemediation(db, remediation); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "La remediation " + title + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "La remediation " + title + " à était créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/remediation-list")
}

func AddSynchronization(c *gin.Context, db *gorm.DB) {

	typ := c.PostForm("title")
	attemptCountStr := c.PostForm("title")
	status := c.PostForm("title")
	startStr := c.PostForm("title")
	endStr := c.PostForm("title")


	attemptCount, err := strconv.Atoi(attemptCountStr)
	if err != nil {
		c.String(http.StatusBadRequest, "attemptCount invalide : %v", err)
		return
	}

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de debut invalide : %v", err)
		return
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de fin invalide : %v", err)
		return
	}

	synchronization := &models.Synchronization{Type: typ, AttemptCount: attemptCount, Status: status, Start: start, End: end}

	if err := models.CreateSynchronization(db, synchronization); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "La synchronization " + typ + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "La synchronization " + typ + " à était créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/synchronization-list")
}

func AddPlanning(c *gin.Context, db *gorm.DB) {

	title := c.PostForm("title")
	typ := c.PostForm("title")
	class := c.PostForm("title")
	schoolYear := c.PostForm("title")
	startStr := c.PostForm("title")
	endStr := c.PostForm("title")


	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de debut invalide : %v", err)
		return
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de fin invalide : %v", err)
		return
	}

	planning := &models.Planning{Title: title, Type: typ, Class: class, SchoolYear: schoolYear, Start: start, End: end}

	if err := models.CreatePlanning(db, planning); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Le planning " + title + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Le planning " + title + " à était créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/planning-list")
}

func AddPlanningSlot(c *gin.Context, db *gorm.DB) {

	planningIDStr := c.PostForm("title")
	lesson := c.PostForm("title")
	teacher := c.PostForm("title")
	startHourStr := c.PostForm("title")
	endHourStr := c.PostForm("title")

	planningID, err := strconv.ParseUint(planningIDStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "planningID invalide : %v", err)
		return
	}

	startHour, err := time.Parse("15:04", startHourStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de debut invalide : %v", err)
		return
	}

	endHour, err := time.Parse("15:04", endHourStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de fin invalide : %v", err)
		return
	}

	planningSlot := &models.PlanningSlot{PlaningID: planningID, Lesson: lesson, Teacher: teacher, StartHour: startHour, EndHour: endHour}

	if err := models.CreatePlanningSlot(db, planningSlot); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "La leçon " + lesson + " du planning n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "La cellule " + lesson + " du planning à était créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/planning-slot-list")
}

func AddSchoolYear(c *gin.Context, db *gorm.DB) {

	title := c.PostForm("title")
	startStr := c.PostForm("startYear")
	endStr := c.PostForm("endYear")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de debut invalide : %v", err)
		return
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de fin invalide : %v", err)
		return
	}

	
	if start.After(end) {
		c.String(http.StatusBadRequest, "Date de debut ne peut pas être supérieure à la date de fin")
		return
	}

	schoolYear := &models.SchoolYear{Title: title, StartYear: start, EndYear: end}

	if err := models.CreateSchoolYear(db, schoolYear); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "L'année scolaire " + title + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "L'année scolaire " + title + " à était créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/setting-school-year")
}

func AddNotification(c *gin.Context, db *gorm.DB) {

	title := c.PostForm("title")
	description := c.PostForm("title")
	status := c.PostForm("title")
	

	notification := &models.Notification{Title: title, Description: description, Status: status}

	if err := models.CreateNotification(db, notification); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "La notification " + title + " n'a pas pu être ajouter"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "La notification " + title + " à était créée avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/notification-list")
}

func AddStudentClass(c *gin.Context, db *gorm.DB) {

	studentIDStr := c.PostForm("title")
	class := c.PostForm("title")
	reason := c.PostForm("title")
	schoolYear := c.PostForm("title")
	
	studentID, err := strconv.ParseUint(studentIDStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "studentID invalide : %v", err)
		return
	}

	studentClass := &models.StudentClass{StudentID: studentID, Class: class, Reason: reason, SchoolYear: schoolYear}

	if err := models.CreateStudentClass(db, studentClass); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "L'apprenant @" + studentIDStr + " n'a pas pu être lier à une class " + class + " avec  succés"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "L'apprenant @" + studentIDStr + " à était lier à une class " + class + " avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/student-class-list")
}

func AddTransactionHistory(c *gin.Context, db *gorm.DB) {

	title := c.PostForm("title")
	priceStr := c.PostForm("title")
	typ := c.PostForm("title")
	actor := c.PostForm("title")
	
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "price invalide : %v", err)
		return
	}

	transactionHistory := &models.TransactionHistory{Title: title, Price: price, Type: typ, Actor: actor}

	if err := models.CreateTransactionHistory(db, transactionHistory); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "La transaction " + title + " n'a pas pu être éffectuer avec  succés"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "La transaction " + title + " a pas pu être éffectuer avec  succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/transaction-history-list")
}

func AddRemuneration(c *gin.Context, db *gorm.DB) {

	userIDStr := c.PostForm("title")
	userType := c.PostForm("title")
	priceStr := c.PostForm("title")
	typ := c.PostForm("title")
	dateStr := c.PostForm("title")

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "userID invalide : %v", err)
		return
	}
	
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "price invalide : %v", err)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	remuneration := &models.Remuneration{UserID: userID, UserType: userType, Price: price, Type: typ, Date: date}

	if err := models.CreateRemuneration(db, remuneration); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "La remunération de l'utilisateur " + userType + ", @" + userIDStr + " n'a pas pu être éffectuer avec succés"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "La remunération de l'utilisateur " + userType + ", @" + userIDStr + " a pas pu être éffectuer avec succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/remuneration-list")
}

func AddLeave(c *gin.Context, db *gorm.DB) {

	userIDStr := c.PostForm("title")
	userType := c.PostForm("title")
	status := c.PostForm("title")
	startYearStr := c.PostForm("title")
	endYearStr := c.PostForm("title")

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "userID invalide : %v", err)
		return
	}

	startYear, err := time.Parse("2006-01-02", startYearStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de debut invalide : %v", err)
		return
	}
	
	endYear, err := time.Parse("2006-01-02", endYearStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de fin invalide : %v", err)
		return
	}	

	
	leave := &models.Leave{UserID: userID, UserType: userType, Status: status, StartYear: startYear, EndYear: endYear}

	if err := models.CreateLeave(db, leave); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Le conger de l'utilisateur " + userType + ", @" + userIDStr + " n'a pas pu être éffectuer avec succés"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "CREATE", Message: "Le conger de l'utilisateur " + userType + ", @" + userIDStr + " a pas pu être éffectuer avec succés"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/leave-list")
}