package main

import (
	"log"
	"net/http"

	"github.com/anaard/simple-student-management/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	
	routes.RegisterStudentManagementRoutes(r)
	
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}