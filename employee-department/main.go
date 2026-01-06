package main

import (
	"fmt"
	"log"

	"github.com/19481A1281/employee-department/config"
	"github.com/19481A1281/employee-department/models"
	"github.com/gin-gonic/gin"
	"github.com/19481A1281/employee-department/repository"
	"github.com/19481A1281/employee-department/services"
	"github.com/19481A1281/employee-department/controllers"
)

func main() {
	

	db := config.ConnectDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database connection")
	}
	defer sqlDB.Close()
	log.Println("Connected to database successfully")
	fmt.Println("Database connection established")

	db.AutoMigrate(&models.Department{},&models.Employee{})

	// 1. Initialize Layers
	deptRepo := repository.NewDepartmentRepository(config.ConnectDB())
	deptService := services.NewDepartmentService(deptRepo)
	deptController := controllers.NewDepartmentController(deptService)

	r := gin.Default()

	// 2. The controller handles its own routing logic
	deptController.RegisterRoutes(r)

	r.Run(":8080")
}