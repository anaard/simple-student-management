package models

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model           // Include following fields to the DB: ID, CreatedAt, UpdatedAt e DeletedAt
	Name         string  `json:"name"`
	ClassId      uint    `gorm:"default:0" json:"class"`
	TotalFaults  uint    `gorm:"default:0" json:"total_faults"`
	GradeAverage float64 `json:"grade_average"`
}

func (Student) TableName() string {
	return "students"
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

func GetStudentbyId(ID int64) (*Student, *gorm.DB) {
	var s Student
	db := db.Where("Id=?", ID).Find(&s)

	return &s, db
}

func DeleteStudent(ID int64) (Student, error) {
	var s Student
	// Find the student with the given ID
	if err := db.First(&s, ID).Error; err != nil {
		// Return an error if the student is not found
		return s, err
	}
	// Delete the student
	if err := db.Delete(&s).Error; err != nil {
		// Return an error if there was an issue deleting the student
		return s, err
	}
	// Return the deleted student
	return s, nil
}
