package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
)

func DeleteClass(w http.ResponseWriter, r *http.Request) { // TODO: Change students classes to None when a class is deleted
	id := getId(r, "classId")

	class, err := models.DeleteClass(id)

	if err != nil {
		writeJSONResponse(w, http.StatusNotFound, models.Class{})
		return
	}

	if class.ID == 0 {
		writeJSONResponse(w, http.StatusInternalServerError, models.Class{})
		return
	}

	writeJSONResponse(w, http.StatusOK, class)
}
