package controllers

import (
	"demo/employee-department/models"
	"demo/employee-department/services"
	"net/http"
	"strconv"

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

func (c *DepartmentController) GetByID(ctx *gin.Context){
	id, _ := strconv.Atoi(ctx.Param("id"))
	
	data, err := c.service.GetByID(uint(id))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "department not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *DepartmentController) UpdateDepartment(ctx *gin.Context){
	id, _ := strconv.Atoi(ctx.Param("id"))

	var departemnt models.Department

	if err := ctx.ShouldBindJSON(&departemnt); err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(), 
		})
		return
	}

	departemnt.ID = uint(id)
	err:= c.service.UpdateDepartment(&departemnt)

	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, departemnt)

}

func (c *DepartmentController) DeleteDepartment(ctx *gin.Context){
	id, _:= strconv.Atoi(ctx.Param("id"))

	err := c.service.DeleteDepartment(uint(id))

	if err!=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Departemnt deleted"})
}