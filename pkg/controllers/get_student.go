package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	students := models.GetAllStudents()
	writeJSONResponse(w, http.StatusOK, students)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id := getStudentId(r)
	
	student, _ := models.GetStudentbyId(id)
	
	writeJSONResponse(w, http.StatusOK, student)
}