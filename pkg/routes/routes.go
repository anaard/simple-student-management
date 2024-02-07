package routes

import (
	"github.com/anaard/simple-student-management/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterStudentManagementRoutes = func(router *mux.Router) {
	//router.HandleFunc("/register/", controllers.Register).Methods("POST")
	//router.HandleFunc("/login/", controllers.Login).Methods("POST")

	router.HandleFunc("/student/", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/student/all", controllers.GetStudents).Methods("GET")

	router.HandleFunc("/student/{studentId}", controllers.GetStudentByID).Methods("GET")
	router.HandleFunc("/student/{studentId}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{studentId}", controllers.DeleteStudent).Methods("DELETE")

	router.HandleFunc("/class/", controllers.CreateClass).Methods("POST")            // create class
	router.HandleFunc("/class/all", controllers.GetClasses).Methods("GET")           // return brief statistics from each class (id, number of students)
	router.HandleFunc("/class/{classId}", controllers.GetClassById).Methods("GET")   // return n° students and the students
	router.HandleFunc("/class/{classId}", controllers.DeleteClass).Methods("DELETE") // return class and remove students from there

	//router.HandleFunc("/{classId}/{studentId}", controllers.EnrollStudent).Methods("POST") // enroll student in class
	//router.HandleFunc("/{classId}/{studentId}", controllers.RemoveStudentFromClass).Methods("DELETE") // remove student from class

}
