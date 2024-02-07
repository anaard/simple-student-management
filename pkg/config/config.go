package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type mySQLConfigData struct {
	Username string `json:"db_username"`
	Password string `json:"db_password"`
	DbName   string `json:"db_name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

var (
	db *gorm.DB
)

func loadmySQLConfigData(filename string) (mySQLConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return mySQLConfigData{}, err
	}

	var config mySQLConfigData
	err = json.Unmarshal(bytes, &config)

	if err != nil {
		return mySQLConfigData{}, err
	}

	return config, nil
}

func Connect() {
	mySQLConfig, err := loadmySQLConfigData("pkg/config/.mySQLconfig")

	if err != nil {
		panic(err)
	}

	dataSource := mySQLConfig.Username + ":" + mySQLConfig.Password + "@tcp(" + mySQLConfig.Host + ":" + mySQLConfig.Port + ")/" + mySQLConfig.DbName + "?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", dataSource)
	
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
