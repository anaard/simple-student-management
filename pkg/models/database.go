package models

import (
	"github.com/anaard/simple-student-management/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// TODO: Ver uma forma de usar herança na class e student (métodos iguais)

func init() { // Function called before main
	config.Connect()

	db = config.GetDB()

	db.AutoMigrate(&Student{}, &Class{})
}
