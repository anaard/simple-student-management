package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func writeJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)
}

func getStudentId(r *http.Request) int64 {
	vars := mux.Vars(r)

	studentId := vars["studentId"]

	id, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("Error while Parsing")
	}
	return id
}
