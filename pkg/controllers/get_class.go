package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
)

func GetClasses(w http.ResponseWriter, r *http.Request) {
	classes := models.GetAllClasses()
	writeJSONResponse(w, http.StatusOK, classes)
}

func GetClassById(w http.ResponseWriter, r *http.Request) {
	id := getId(r, "classId")

	class, _ := models.GetClassById(id)

	writeJSONResponse(w, http.StatusOK, class)
}
