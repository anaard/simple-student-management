package models

import (
	"github.com/jinzhu/gorm"
)

type Class struct {
	gorm.Model
	TotalStudentNumber uint      `gorm:"default:0"`
	Students           []Student `gorm:"foreignKey:ClassId"`
}

func (Class) TableName() string {
	return "classes"
}

func (class *Class) CreateClass() *Class {
	db.NewRecord(class)
	db.Create(&class)
	return class
}

func GetClassById(ID int64) (*Class, *gorm.DB) {
	var class Class
	db := db.Where("Id=?", ID).Find(&class)

	return &class, db
}

func DeleteClass(ID int64) (Class, error) {
	var class Class

	if err := db.First(&class, ID).Error; err != nil {
		return class, err
	}
	// Delete the student
	if err := db.Delete(&class).Error; err != nil {
		return class, err
	}
	return class, nil
}

func GetAllStudentsInClass(classID uint) ([]Student, error) {
	var students []Student
	if err := db.Where("class_id = ?", classID).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func GetAllClasses() []Class {
	var Classes []Class
	db.Find(&Classes)
	return Classes
}
