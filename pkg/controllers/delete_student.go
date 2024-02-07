package controllers

// TODO: Ver status de erros corretos
import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
)

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := getId(r, "studentId")

	student, err := models.DeleteStudent(id)

	if err != nil {
		writeJSONResponse(w, http.StatusNotFound, models.Student{})
		return
	}

	if student.ID == 0 {
		writeJSONResponse(w, http.StatusInternalServerError, models.Student{})
		return
	}

	writeJSONResponse(w, http.StatusOK, student)
}
