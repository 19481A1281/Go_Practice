package controllers

import (
	"net/http"
	"demo/employee-department/models"
	"demo/employee-department/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	service services.EmployeeService
}

func NewEmployeeController(service services.EmployeeService) *EmployeeController {
	return &EmployeeController{
		service: service,
	}
}

func (c *EmployeeController) Create(ctx *gin.Context) {

	var employee models.Employee

	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := c.service.CreateEmployee(&employee)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, employee)
}

func (c *EmployeeController) GetAll(ctx *gin.Context) {

	data, err := c.service.GetEmployees()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *EmployeeController) GetByID(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := c.service.GetEmployeeByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "employee not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}

func (c *EmployeeController) Delete(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.service.DeleteEmployee(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "employee deleted",
	})
}