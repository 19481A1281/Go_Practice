package routes

import (
	"demo/employee-department/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	employeeController *controllers.EmployeeController,
	departmentController *controllers.DepartmentController,
) {

	employees := router.Group("/employees")
	{
		employees.POST("/", employeeController.Create)
		employees.GET("/", employeeController.GetAll)
		employees.GET("/:id", employeeController.GetByID)
		employees.DELETE("/:id", employeeController.Delete)
		//employees.PUT("/:id",employeeController.Update)
		employees.PATCH("/:id",employeeController.Update)
	}

	departments := router.Group("/departments")
	{
		departments.POST("/", departmentController.Create)
		departments.GET("/", departmentController.GetAll)
		departments.GET("/:id",departmentController.GetByID)
		departments.PUT("/:id",departmentController.UpdateDepartment)
		departments.DELETE("/:id",departmentController.DeleteDepartment)
	}
}