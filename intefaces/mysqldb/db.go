package mysqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Mysql struct{
	db *sql.DB
}

func New(dbUser, dbPassword, dbHost, dbPort, dbName string) (*Mysql, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Println(dsn)
	db , err := sql.Open("mysql", dsn)

	if err!=nil{
		log.Println("connection failed:",err)
		return nil,err
	}

	err = db.Ping()
	if err !=nil{
		log.Println("mysql ping failed:",err)
		return nil,err
	}

	return &Mysql{db:db},nil
}

func (mySql Mysql) Close(){
	err := mySql.db.Close()

	if err!=nil{
		log.Println("error in closing in the connection:",err)
	}
}

func (mySql Mysql) InsertUser(userName string) error{
	mySql.db.Exec("INSERT...")

	return nil
}

func (Mysql Mysql) SelectUser(userName string) (string,error){
	Mysql.db.Exec("SELECT...")

	return "user",nil
}