package controllers

// TODO: delete returning empty
import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
)

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := getStudentId(r)

	student := models.DeleteStudent(id)

	writeJSONResponse(w, http.StatusOK, student)
}
