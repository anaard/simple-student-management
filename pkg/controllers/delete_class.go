package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func (s SystemController) DeleteClass(w http.ResponseWriter, r *http.Request) { // TODO: Change students classes to None when a class is deleted
	id := utils.GetId(r, "classId")

	class, err := models.DeleteClass(id)

	if err != nil {
		utils.WriteJSONResponse(w, http.StatusNotFound, models.Class{})
		return
	}

	if class.ID == 0 {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, models.Class{})
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, class)
}
