package models

import (
	"github.com/anaard/simple-student-management/pkg/config"
	"github.com/jinzhu/gorm"
	//"time"
)

var db *gorm.DB
// TODO: Incluir campos: Total de Faltas, média de notas (score)
type Student struct {
	gorm.Model // Include following fields to the DB: ID, CreatedAt, UpdatedAt e DeletedAt
	Name       string `json:"name"`
	Birth      string `json:"birth"` // Change to time
	Class      string `json:"class"`
}

func init() { // Function called before main
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
}

func (s *Student) CreateStudent() *Student {
	db.NewRecord(s)
	db.Create(&s)
	return s
}

func GetAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}

func GetStudentbyId(Id int64) (*Student, *gorm.DB) {
	var s Student
	db := db.Where("Id=?", Id).Find(&s)
	return &s, db
}

func DeleteStudent(ID int64) Student {
	var s Student
	db.First(&s, ID)
	db.Delete(&s)
	return s
}