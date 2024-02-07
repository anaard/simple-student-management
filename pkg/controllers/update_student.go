package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var updateStudent = &models.Student{}
	utils.ParseBody(r, updateStudent)
	
	id := getStudentId(r)
	
	student, db := models.GetStudentbyId(id)
	
	if updateStudent.Name != "" {
		student.Name = updateStudent.Name
	}
	
	if updateStudent.Birth != "" {
		student.Birth = updateStudent.Birth
	}
	
	if updateStudent.Class != "" {
		student.Class = updateStudent.Class
	}
	
	db.Save(&student) // TODO: Handle this error
	writeJSONResponse(w, http.StatusOK, student)
}