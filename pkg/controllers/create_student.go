package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	createStudent := &models.Student{}
	utils.ParseBody(r, createStudent)

	student := createStudent.CreateStudent()

	writeJSONResponse(w, http.StatusOK, student)
}
