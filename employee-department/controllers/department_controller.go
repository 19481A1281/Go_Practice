package controllers

import (
	"demo/employee-department/models"
	"demo/employee-department/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DepartmentController struct{
	service services.DepartmentService
}

func NewDepartmentController(service services.DepartmentService) *DepartmentController{
	return &DepartmentController{
		service: service,
	}
}

func(c *DepartmentController) Create(ctx *gin.Context){
	var department models.Department

	if err := ctx.ShouldBindJSON(&department); err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := c.service.CreateDepartment(&department)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, department)

}

func (c *DepartmentController) GetAll(ctx *gin.Context){
	data, err := c.service.GetDepartments()

	if err!=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}