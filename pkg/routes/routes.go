package routes

import (
	"github.com/gorilla/mux"
	"github.com/anaard/simple-student-management/pkg/controllers"
)

var RegisterStudentManagementRoutes = func(router *mux.Router) {
	router.HandleFunc("/student/", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/student/all", controllers.GetStudents).Methods("GET")
	router.HandleFunc("/student/{studentId}", controllers.GetStudentByID).Methods("GET")
	router.HandleFunc("/student/{studentId}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{studentId}", controllers.DeleteStudent).Methods("DELETE")
}