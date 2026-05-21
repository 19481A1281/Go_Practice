package main

import (
	"demo/employee-department/config"
	"demo/employee-department/controllers"
	"demo/employee-department/models"
	"demo/employee-department/repositories"
	"demo/employee-department/routes"
	"demo/employee-department/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("Welcome to employee - department application")

	mysqlCon, err := config.NewMysqlConnection()
	if err !=nil{
		panic(err)
	}

	mysqlCon.DB.AutoMigrate(
		&models.Department{},
		&models.Employee{},
	)

	employeeRepo := repositories.NewEmployeeRepository(mysqlCon.DB)
	departmentRepo := repositories.NewDepartmentRepositoty(mysqlCon.DB)

	employeeService := services.NewEmployeeService(employeeRepo)
	departmentService := services.NewDepartmentService(departmentRepo)


	employeeController := controllers.NewEmployeeController(employeeService)
	departmentController := controllers.NewDepartmentController(departmentService)

	router := gin.Default()
	routes.SetupRoutes(
		router,
		employeeController,
		departmentController,
	)

	router.Run(":8080")
}