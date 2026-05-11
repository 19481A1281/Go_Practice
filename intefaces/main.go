package main

import (
	"fmt"
	"log"

	"example.com/dbconnect/mysqldb"
)

type dbConnect interface{
	Close()
	InsertUser(userName string) error
	SelectUser(userName string) (string, error)
}

type Application struct{
	db dbConnect
}

func NewApplication(db dbConnect) *Application{
	return &Application{db:db}
}

func(app Application) Run(){
	fmt.Println("App is running")
}

func main(){
	db,err := mysqldb.New("root","root","localhost","3306","xyz")

	if err!=nil{
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer db.Close()

	app := NewApplication(db)
	app.Run()
}