package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var updateStudent = &models.Student{}
	utils.ParseBody(r, updateStudent)

	id := getId(r, "studentId")

	student, db := models.GetStudentbyId(id)

	if student.ID == 0 { // Student doesn't exist on the db
		writeJSONResponse(w, http.StatusInternalServerError, models.Student{})
		return
	}

	if updateStudent.Name != "" {
		student.Name = updateStudent.Name
	}

	if updateStudent.ClassId > 0 {
		student.ClassId = updateStudent.ClassId
	}

	db.Save(&student) // TODO: Handle this error
	writeJSONResponse(w, http.StatusOK, student)
}
