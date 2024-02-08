package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func (s SystemController) CreateStudent(w http.ResponseWriter, r *http.Request) {
	createStudent := &models.Student{}
	utils.ParseBody(r, createStudent)

	student := createStudent.CreateStudent()

	utils.WriteJSONResponse(w, http.StatusOK, student)
}
