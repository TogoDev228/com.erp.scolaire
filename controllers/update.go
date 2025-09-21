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
	id, err := getIDParam(c)
	
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
	endStr := c.PostForm("end")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	staff := &models.Staff{}
	if err := db.First(staff, id).Error; err != nil {
		c.String(http.StatusNotFound, "Employé(e) non trouvé(e) : %v", err)
		return
	}

	if (start.After(end)){
		c.String(http.StatusBadRequest, "La date de début ne peut pas être postérieure à la date de résilition")
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
	staff.End = end

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
	valueStr := c.PostForm("value")
	quantityStr := c.PostForm("quantity")
	statut := c.PostForm("statut")
	startStr := c.PostForm("start")
	repairStr := c.PostForm("repair")

	value, err := strconv.Atoi(valueStr)

	if err != nil {
		c.String(http.StatusBadRequest, "Value invalide : %v", err)
		return
	}

	// Conversion de start en time.Time
	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	repair, err := time.Parse("2006-01-02", repairStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date invalide : %v", err)
		return
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Quantité invalide : %v", err)
		return
	}

	if (start.After(repair)){
		c.String(http.StatusBadRequest, "La date de mise en service ne peut pas être postérieure à la date de dernière réparation")
		return
	}

	// Récupérer l'item existant depuis la BDD
	item, err := models.GetItemByID(db, uint64(id))
	if err != nil {
		c.String(http.StatusNotFound, "Ressource non trouvé : %v", err)
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
	item.Repair = repair

	// Enregistrer la mise à jour en BDD
	if err := models.UpdateItem(db, item); err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la mise à jour : %v", err)

		log := &models.Log{Type: "ERROR", Message: "Ressource " + title + " n'a pas pu être mis à jour"}
		if err := models.CreateLog(db, log); err != nil {
			c.String(http.StatusInternalServerError, "Erreur lors de la création du log : %v", err)
			return
		}
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "Ressource " + title + " mis à jour avec succès"}
	if err := models.CreateLog(db, log); err != nil {
		c.String(http.StatusInternalServerError, "Erreur lors de la création du log : %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/item-list")
}

func UpdatePayment(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	var payment models.Payment
	if err := db.First(&payment, id).Error; err != nil {
		c.String(http.StatusNotFound, "Payment introuvable : %v", err)
		return
	}

	studentIDStr := c.PostForm("studentID")
	priceStr := c.PostForm("price")
	typ := c.PostForm("type")

	studentID, _ := strconv.ParseUint(studentIDStr, 10, 64)
	price, _ := strconv.ParseFloat(priceStr, 64)

	payment.StudentID = studentID
	payment.Price = price
	payment.Type = typ

	if err := models.UpdatePayment(db, &payment); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: "Le paiement #" + studentIDStr + " n'a pas pu être mis à jour"})
		return
	}

	models.CreateLog(db, &models.Log{Type: "UPDATE", Message: "Le paiement #" + studentIDStr + " a été mis à jour avec succès"})
	c.Redirect(http.StatusSeeOther, "/payment-list")
}

func UpdateTeacherLesson(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	var teacherLesson models.TeacherLesson
	if err := db.First(&teacherLesson, id).Error; err != nil {
		c.String(http.StatusNotFound, "Relation introuvable : %v", err)
		return
	}

	teacherIDStr := c.PostForm("studentID")
	lesson := c.PostForm("price")
	schoolYear := c.PostForm("type")

	teacherID, _ := strconv.ParseUint(teacherIDStr, 10, 64)

	teacherLesson.TeacherID = teacherID
	teacherLesson.Lesson = lesson
	teacherLesson.SchoolYear = schoolYear

	if err := models.UpdateTeacherLesson(db, &teacherLesson); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: "La relation professeur @" + teacherIDStr + " - " + lesson + " n'a pas pu être mise à jour"})
		return
	}

	models.CreateLog(db, &models.Log{Type: "UPDATE", Message: "La relation professeur @" + teacherIDStr + " - " + lesson + " a été mise à jour avec succès"})
	c.Redirect(http.StatusSeeOther, "/teacher-lesson-list")
}

func UpdateAttendance(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	var attendance models.Attendance
	if err := db.First(&attendance, id).Error; err != nil {
		c.String(http.StatusNotFound, "Présence introuvable : %v", err)
		return
	}

	userIDStr := c.PostForm("studentID")
	userType := c.PostForm("price")
	typ := c.PostForm("price")
	dateStr := c.PostForm("price")
	schoolYear := c.PostForm("type")

	userID, _ := strconv.ParseUint(userIDStr, 10, 64)
	date, _ := time.Parse("2006-01-02", dateStr)

	attendance.UserID = userID
	attendance.UserType = userType
	attendance.Type = typ
	attendance.Date = date
	attendance.SchoolYear = schoolYear

	if err := models.UpdateAttendance(db, &attendance); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: "La présence de l'utilisateur @" + userIDStr + " n'a pas pu être mise à jour"})
		return
	}

	models.CreateLog(db, &models.Log{Type: "UPDATE", Message: "La présence de l'utilisateur @" + userIDStr + " a été mise à jour avec succès"})
	c.Redirect(http.StatusSeeOther, "/attendance-list")
}

func UpdateRole(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	var role models.Role
	if err := db.First(&role, id).Error; err != nil {
		c.String(http.StatusNotFound, "Rôle introuvable : %v", err)
		return
	}

	title := c.PostForm("title")
	role.Title = title

	if err := models.UpdateRole(db, &role); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: "Le rôle " + title + " n'a pas pu être mis à jour"})
		return
	}

	models.CreateLog(db, &models.Log{Type: "UPDATE", Message: "Le rôle " + title + " a été mis à jour avec succès"})
	c.Redirect(http.StatusSeeOther, "/setting-role")
}

func UpdateExpense(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	var expense models.Expense
	if err := db.First(&expense, id).Error; err != nil {
		c.String(http.StatusNotFound, "Dépense introuvable : %v", err)
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("title")
	typ := c.PostForm("title")
	status := c.PostForm("title")
	schoolYear := c.PostForm("title")
	creator := c.PostForm("title")
	startStr := c.PostForm("title")
	endStr := c.PostForm("title")

	start, _ := time.Parse("2006-01-02", startStr)
	end, _ := time.Parse("2006-01-02", endStr)

	expense.Title = title
	expense.Description = description
	expense.Type = typ
	expense.Status = status
	expense.SchoolYear = schoolYear
	expense.Creator = creator
	expense.Start = start
	expense.End = end

	if err := models.UpdateExpense(db, &expense); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: "La dépense " + title + " n'a pas pu être mise à jour"})
		return
	}

	models.CreateLog(db, &models.Log{Type: "UPDATE", Message: "La dépense " + title + " a été mise à jour avec succès"})
	c.Redirect(http.StatusSeeOther, "/expense-list")
}

func UpdateTask(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		c.String(http.StatusNotFound, "Tâche introuvable : %v", err)
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("title")
	typ := c.PostForm("title")
	status := c.PostForm("title")
	schoolYear := c.PostForm("title")
	creator := c.PostForm("title")
	startStr := c.PostForm("title")
	endStr := c.PostForm("title")

	start, _ := time.Parse("2006-01-02", startStr)
	end, _ := time.Parse("2006-01-02", endStr)

	task.Title = title
	task.Description = description
	task.Type = typ
	task.Status = status
	task.SchoolYear = schoolYear
	task.Creator = creator
	task.Start = start
	task.End = end

	if err := models.UpdateTask(db, &task); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: "La tâche " + title + " n'a pas pu être mise à jour"})
		return
	}

	models.CreateLog(db, &models.Log{Type: "UPDATE", Message: "La tâche " + title + " a été mise à jour avec succès"})
	c.Redirect(http.StatusSeeOther, "/task-list")
}

func UpdateExam(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	var exam models.Exam
	if err := db.First(&exam, id).Error; err != nil {
		c.String(http.StatusNotFound, "Examen introuvable : %v", err)
		return
	}

	title := c.PostForm("title")
	typ := c.PostForm("title")
	status := c.PostForm("title")
	schoolYear := c.PostForm("title")
	startStr := c.PostForm("title")
	endStr := c.PostForm("title")

	start, _ := time.Parse("2006-01-02", startStr)
	end, _ := time.Parse("2006-01-02", endStr)

	exam.Title = title
	exam.Type = typ
	exam.Status = status
	exam.SchoolYear = schoolYear
	exam.Start = start
	exam.End = end

	if err := models.UpdateExam(db, &exam); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		models.CreateLog(db, &models.Log{Type: "ERROR", Message: "L'examen " + title + " n'a pas pu être mis à jour"})
		return
	}

	models.CreateLog(db, &models.Log{Type: "UPDATE", Message: "L'examen " + title + " a été mis à jour avec succès"})
	c.Redirect(http.StatusSeeOther, "/exam-list")
}

func UpdateRemediation(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	title := c.PostForm("title")
	description := c.PostForm("description")
	studentIDStr := c.PostForm("studentID")
	teacherIDStr := c.PostForm("teacherID")
	status := c.PostForm("status")
	startStr := c.PostForm("start")
	endStr := c.PostForm("end")

	studentID, _ := strconv.ParseUint(studentIDStr, 10, 64)
	teacherID, _ := strconv.ParseUint(teacherIDStr, 10, 64)
	start, _ := time.Parse("2006-01-02", startStr)
	end, _ := time.Parse("2006-01-02", endStr)

	remediation := &models.Remediation{
		ID:          id,
		Title:       title,
		Description: description,
		StudentID:   studentID,
		TeacherID:   teacherID,
		Status:      status,
		Start:       start,
		End:         end,
	}

	if err := models.UpdateRemediation(db, remediation); err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		log := &models.Log{Type: "ERROR", Message: "La remediation " + title + " n'a pas pu être mise à jour"}
		models.CreateLog(db, log)
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "La remediation " + title + " a été mise à jour avec succès"}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/remediation-list")
}

func UpdatePlanning(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	title := c.PostForm("title")
	typ := c.PostForm("type")
	class := c.PostForm("class")
	schoolYear := c.PostForm("school_year")
	startStr := c.PostForm("start")
	endStr := c.PostForm("end")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de début invalide : %v", err)
		return
	}
	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de fin invalide : %v", err)
		return
	}

	updates := map[string]interface{}{
		"Title":      title,
		"Type":       typ,
		"Class":      class,
		"SchoolYear": schoolYear,
		"Start":      start,
		"End":        end,
	}

	if err := db.Model(&models.Planning{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "Le planning a été mis à jour avec succès"}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/planning-list")
}

func UpdatePlanningSlot(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	planningIDStr := c.PostForm("planning_id")
	lesson := c.PostForm("lesson")
	teacher := c.PostForm("teacher")
	startHourStr := c.PostForm("start_hour")
	endHourStr := c.PostForm("end_hour")

	planningID, err := strconv.ParseUint(planningIDStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "planningID invalide : %v", err)
		return
	}
	startHour, err := time.Parse("15:04", startHourStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Heure de début invalide : %v", err)
		return
	}
	endHour, err := time.Parse("15:04", endHourStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Heure de fin invalide : %v", err)
		return
	}

	updates := map[string]interface{}{
		"PlaningID": planningID,
		"Lesson":    lesson,
		"Teacher":   teacher,
		"StartHour": startHour,
		"EndHour":   endHour,
	}

	if err := db.Model(&models.PlanningSlot{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "La cellule du planning a été mise à jour avec succès"}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/planning-slot-list")
}

func UpdateSchoolYear(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	title := c.PostForm("title")
	startStr := c.PostForm("startYear")
	endStr := c.PostForm("endYear")

	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de début invalide : %v", err)
		return
	}
	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de fin invalide : %v", err)
		return
	}

	if start.After(end) {
		c.String(http.StatusBadRequest, "Date de debut ne peut être supèrieur à celui de la fin")
		return
	}

	updates := map[string]interface{}{
		"Title":     title,
		"StartYear": start,
		"EndYear":   end,
	}

	if err := db.Model(&models.SchoolYear{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "L'année scolaire a été mise à jour avec succès"}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/setting-school-year")
}

func UpdateStudentClass(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	studentIDStr := c.PostForm("student_id")
	class := c.PostForm("class")
	reason := c.PostForm("reason")
	schoolYear := c.PostForm("school_year")

	studentID, err := strconv.ParseUint(studentIDStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "studentID invalide : %v", err)
		return
	}

	updates := map[string]interface{}{
		"StudentID":  studentID,
		"Class":      class,
		"Reason":     reason,
		"SchoolYear": schoolYear,
	}

	if err := db.Model(&models.StudentClass{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "La liaison de l'étudiant a été mise à jour avec succès"}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/student-class-list")
}

func UpdateRemuneration(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	userIDStr := c.PostForm("user_id")
	userType := c.PostForm("user_type")
	priceStr := c.PostForm("price")
	typ := c.PostForm("type")
	dateStr := c.PostForm("date")

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

	updates := map[string]interface{}{
		"UserID":   userID,
		"UserType": userType,
		"Price":    price,
		"Type":     typ,
		"Date":     date,
	}

	if err := db.Model(&models.Remuneration{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "La rémunération a été mise à jour avec succès"}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/remuneration-list")
}

func UpdateLeave(c *gin.Context, db *gorm.DB) {
	id, err := getIDParam(c)
	if err != nil {
		c.String(http.StatusBadRequest, "ID invalide : %v", err)
		return
	}

	userIDStr := c.PostForm("user_id")
	userType := c.PostForm("user_type")
	status := c.PostForm("status")
	startYearStr := c.PostForm("start_year")
	endYearStr := c.PostForm("end_year")

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "userID invalide : %v", err)
		return
	}
	startYear, err := time.Parse("2006-01-02", startYearStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de début invalide : %v", err)
		return
	}
	endYear, err := time.Parse("2006-01-02", endYearStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Date de fin invalide : %v", err)
		return
	}

	updates := map[string]interface{}{
		"UserID":    userID,
		"UserType":  userType,
		"Status":    status,
		"StartYear": startYear,
		"EndYear":   endYear,
	}

	if err := db.Model(&models.Leave{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		c.String(http.StatusInternalServerError, "Erreur : %v", err)
		return
	}

	log := &models.Log{Type: "UPDATE", Message: "Le congé a été mis à jour avec succès"}
	models.CreateLog(db, log)

	c.Redirect(http.StatusSeeOther, "/leave-list")
}
