package controllers

// TODO: Ver status de erros corretos
import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func (s SystemController) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(r, "studentId")

	student, err := models.DeleteStudent(id)

	if err != nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, models.Student{})
		return
	}

	if student.ID == 0 {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, models.Student{})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, student)
}
