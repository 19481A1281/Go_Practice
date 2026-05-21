package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDb struct{
	DB *gorm.DB
}


func NewMysqlConnection() (*MysqlDb, error){
	username := "root"
	password := "root"
	host := "127.0.0.1"
	port := "3306"
	dbName := "company_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		username, password, host, port, dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil{
		return nil,err
	}

	db = db.Debug()

	fmt.Println("Databse connected successfully")
	return &MysqlDb{DB:db},nil

}
