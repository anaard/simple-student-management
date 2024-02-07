package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func CreateClass(w http.ResponseWriter, r *http.Request) {
	createClass := &models.Class{}
	utils.ParseBody(r, createClass)

	class := createClass.CreateClass()

	writeJSONResponse(w, http.StatusOK, class)
}
