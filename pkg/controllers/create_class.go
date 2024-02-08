package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func (s SystemController) CreateClass(w http.ResponseWriter, r *http.Request) {
	createClass := &models.Class{}
	utils.ParseBody(r, createClass)

	class := createClass.CreateClass()

	utils.WriteJSONResponse(w, http.StatusOK, class)
}
